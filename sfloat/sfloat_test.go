package sfloat

import (
	"testing"

	"github.com/shawnfeng/sutil2/slog"
)

func TestSFloat(t *testing.T) {

	sf := NewSFloatBySet(8, 5)

	slog.Infof("sfloat:%s", sf)

	if sf.GetPrecNum() != 8 {
		t.Errorf("prec num fail")
	}

	if sf.GetScaleNum() != 5 {
		t.Errorf("scaleNum fail")
	}

	if sf.GetPrecision() != 1e-8 {
		t.Errorf("prec fail")

	}

	if sf.GetScaleFactor() != 1e5 {
		t.Errorf("scale fail")

	}

}
