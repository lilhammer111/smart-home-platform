package product_service

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/binding"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
	"testing"
)

func TestRateProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRateProductService(ctx)
	// init req and assert value

	req := &product.RatingReq{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}

func BenchmarkDatabaseMethod(b *testing.B) {
	binding.SetConfPath("", "../../../conf/dev_conf.yaml")
	binding.SetEnvPath("../../../conf/.env")
	for i := 0; i < b.N; i++ {
		// 调用使用数据库方法的函数
		ctx := context.Background()
		s := NewRateProductService(ctx)
		// init req and assert value

		req := &product.RatingReq{
			ProductId: 1,
			Rating:    4.3,
		}
		_, err := s.Run(req)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGoMethod(b *testing.B) {
	binding.SetConfPath("", "../../../conf/dev_conf.yaml")
	binding.SetEnvPath("../../../conf/.env")
	for i := 0; i < b.N; i++ {
		// 调用使用 Go 计算的函数
		ctx := context.Background()
		s := NewRateProductService(ctx)
		// init req and assert value

		req := &product.RatingReq{
			ProductId: 1,
			Rating:    4.3,
		}
		_, err := s.RunTest(req)
		if err != nil {
			panic(err)
		}
	}
}

// calculate in the database
//goos: linux
//goarch: amd64
//pkg: git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/product_service
//cpu: Intel(R) Xeon(R) CPU E5-2660 0 @ 2.20GHz
//BenchmarkDatabaseMethod
//BenchmarkDatabaseMethod-16    	      69	  20137887 ns/op
//PASS

// calculate in go program
//goos: linux
//goarch: amd64
//pkg: git.zqbjj.top/pet/services/cmd/rpc/product/biz/service/product_service
//cpu: Intel(R) Xeon(R) CPU E5-2660 0 @ 2.20GHz
//BenchmarkGoMethod
//BenchmarkGoMethod-16    	      75	  17513725 ns/op
//PASS
