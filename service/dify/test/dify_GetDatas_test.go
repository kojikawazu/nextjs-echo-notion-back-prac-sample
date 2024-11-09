package test_service_dify

import (
	"testing"

	"github.com/stretchr/testify/assert"

	model_dify "github.com/kojikawazu/backend/model/dify"
	service_dify "github.com/kojikawazu/backend/service/dify"
)

// 正常系のテスト
func TestGetDatas_Success(t *testing.T) {
	// モックサービスの作成
	mockService := new(service_dify.MockDifyService)

	// テストデータの準備
	inputs := map[string]interface{}{
		"message": "テストメッセージ",
	}
	expectedResponse := &model_dify.DifyResponse{
		Title:    "テストタイトル",
		Contents: "テスト内容",
		Kind:     "テスト種類",
	}

	// モックの振る舞いを設定
	mockService.On("GetDatas", inputs).Return(expectedResponse, nil)

	// テスト実行
	result, err := mockService.GetDatas(inputs)

	// アサーション
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	assert.Equal(t, "テストタイトル", result.Title)
	assert.Equal(t, "テスト内容", result.Contents)
	assert.Equal(t, "テスト種類", result.Kind)

	// モックの呼び出しを検証
	mockService.AssertExpectations(t)
}

// エラー系のテスト
func TestGetDatas_Error(t *testing.T) {
	// モックサービスの作成
	mockService := new(service_dify.MockDifyService)

	// テストデータの準備
	inputs := map[string]interface{}{
		"message": "エラーメッセージ",
	}

	// モックの振る舞いを設定
	mockService.On("GetDatas", inputs).Return(nil, assert.AnError)

	// テスト実行
	result, err := mockService.GetDatas(inputs)

	// アサーション
	assert.Error(t, err)
	assert.Nil(t, result)

	// モックの呼び出しを検証
	mockService.AssertExpectations(t)
}
