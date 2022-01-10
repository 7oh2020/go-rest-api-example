package common

import "time"

// 現在時刻を取得します。
// この関数を仕様することでフォーマットやタイムゾーンなどをアプリケーション側で一元管理できます。
func CurrentTime() time.Time {
	return time.Now()
}
