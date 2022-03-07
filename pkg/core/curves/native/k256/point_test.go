package k256_test

import (
	"github.com/dB2510/kryptology/pkg/core/curves/native"
	"github.com/dB2510/kryptology/pkg/core/curves/native/k256"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestK256PointArithmetic_Hash(t *testing.T) {
	var b [32]byte
	sc, err := k256.K256PointNew().Hash(b[:], native.EllipticPointHasherSha256())

	require.NoError(t, err)
	require.True(t, !sc.IsIdentity())
	require.True(t, sc.IsOnCurve())
}
