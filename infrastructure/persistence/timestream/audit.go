package timestream

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"grafana-api/domain/model"
	"grafana-api/infrastructure/secretmanager"
	"strconv"
	"time"
)

type IAuditRepository interface {
	SaveAudit(ctx context.Context, auditDTO model.AuditDTO) error
}

type auditRepository struct {
	TableName string
	DBName    string
}

func NewAuditRepository(secret secretmanager.DBSecret) IAuditRepository {
	return &auditRepository{
		TableName: secret.TableName,
		DBName:    secret.DBName,
	}
}

func (r *auditRepository) SaveAudit(ctx context.Context, auditDTO model.AuditDTO) error {
	sess, err := session.NewSession()
	if err != nil {
		return fmt.Errorf("session error: %w", err)
	}

	client := timestreamwrite.New(sess)
	record := &timestreamwrite.WriteRecordsInput{
		DatabaseName: aws.String(r.DBName),
		TableName:    aws.String(r.TableName),
		Records: []*timestreamwrite.Record{{
			Dimensions: []*timestreamwrite.Dimension{
				{
					Name:  aws.String("Method"),
					Value: aws.String(auditDTO.Method),
				},
				{
					Name:  aws.String("Path"),
					Value: aws.String(auditDTO.Path),
				},
			},
			MeasureName:      aws.String("Payload"),
			MeasureValue:     aws.String(auditDTO.Payload),
			MeasureValueType: aws.String("VARCHAR"),
			Time:             aws.String(strconv.FormatInt(time.Now().Unix(), 10)),
			TimeUnit:         aws.String("SECONDS"),
		}},
	}

	if _, err = client.WriteRecordsWithContext(ctx, record); err != nil {
		return fmt.Errorf("write audit error: %w", err)
	}

	return nil
}
