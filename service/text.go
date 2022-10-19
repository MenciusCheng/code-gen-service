package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MenciusCheng/code-gen-service/model"
	"strings"
	"text/template"
)

func (s *Service) TextGen(ctx context.Context, req model.TextGenReq) (res model.TextGenRes, err error) {
	data, err := JsonStringToMap(req.JsonString)
	if err != nil {
		err = fmt.Errorf("JsonStringToMap err: %w", err)
		return
	}

	text, err := MapGenText(data, req.Template)
	if err != nil {
		err = fmt.Errorf("MapGenText err: %w", err)
		return
	}

	res.Text = text
	return
}

func JsonStringToMap(jsonString string) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &res)
	if err != nil {
		return res, fmt.Errorf("json Unmarshal err: %w", err)
	}
	return res, nil
}

func MapGenText(data map[string]interface{}, tmplStr string) (string, error) {
	tmpl, err := template.New("").Parse(tmplStr)
	if err != nil {
		return "", fmt.Errorf("template Parse err: %w", err)
	}
	// strings.Builder 转化字符串时比 bytes.Buffer 更快
	sb := strings.Builder{}
	err = tmpl.Execute(&sb, data)
	if err != nil {
		return "", fmt.Errorf("tmpl Execute err: %w", err)
	}
	return sb.String(), nil
}
