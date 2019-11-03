/*
 * @Descripttion:
 * @version:
 * @Author: renjie.zhang
 * @Date: 2019-11-03 10:48:57
 * @LastEditTime: 2019-11-03 10:49:20
 */
package Testing

import (
	"testing"
)

func TestIntMin(t *testing.T) {
	ans := IntMin(9, 10)
	if ans != 9 {
		t.Errorf("IntMin(9, 10) = %d; want 9", ans)
	}
}
