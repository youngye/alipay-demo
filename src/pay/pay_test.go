package pay

import (
	//go库包
	"log"
	"testing"
	//"time"
	"net/url"
)


func TestPay(t *testing.T) {
	sStr	:="_input_charset=utf-8&format=xml&partner=2088801631363292&req_data=<direct_trade_create_req><notify_url>http://www2.hunliji.com:8080/alipay/notify_url.jsp</notify_url><call_back_url>http://www2.hunliji.com:8080/alipay/call_back_url.jsp</call_back_url><seller_account_name>firecloud168@gmail.com</seller_account_name><out_trade_no>100001</out_trade_no><subject>测试订单(1分钱)</subject><total_fee>0.01</total_fee><merchant_url>http://www2.hunliji.com:8080/alipay/index.jsp</merchant_url></direct_trade_create_req>&req_id=20141016131529&sec_id=MD5&service=alipay.wap.trade.create.direct&v=2.0"
	sSign	:=alipayMd5Sign(sStr)
	log.Printf("sSign: %s", sSign)
	
	sStr	="out_trade_no=100004\u0026request_token=requestToken\u0026result=success\u0026trade_no=2014101634996227\u0026sign=f810cbddcc39493389bf57a534d3b11d\u0026sign_type=MD5"
	vals, _	:=url.ParseQuery(sStr)
	vals.Del("sign")
	vals.Del("sign_type")
	sStr	=vals.Encode()
	sSign	=alipayMd5Sign(sStr)
	log.Println(vals)
	log.Print(sStr)
	log.Printf("Sign:%s", sSign)
	
	//gotoSign_params.append("service="+params.get("service"));
 //   	gotoSign_params.append("&v="+params.get("v"));
 //   	gotoSign_params.append("&sec_id="+params.get("sec_id"));
 //   	gotoSign_params.append("&notify_data="+params.get("notify_data"));
	
	sStr	="service=alipay.wap.trade.create.direct&v=1.0&sec_id=MD5&notify_data=%3Cnotify%3E%3Cpayment_type%3E1%3C%2Fpayment_type%3E%3Csubject%3E%E6%B5%8B%E8%AF%95%E5%95%86%E5%93%81%3C%2Fsubject%3E%3Ctrade_no%3E2014101634996227%3C%2Ftrade_no%3E%3Cbuyer_email%3Eyycmail%40163.com%3C%2Fbuyer_email%3E%3Cgmt_create%3E2014-10-16+14%3A57%3A01%3C%2Fgmt_create%3E%3Cnotify_type%3Etrade_status_sync%3C%2Fnotify_type%3E%3Cquantity%3E1%3C%2Fquantity%3E%3Cout_trade_no%3E100004%3C%2Fout_trade_no%3E%3Cnotify_time%3E2014-10-16+14%3A57%3A13%3C%2Fnotify_time%3E%3Cseller_id%3E2088801631363292%3C%2Fseller_id%3E%3Ctrade_status%3ETRADE_SUCCESS%3C%2Ftrade_status%3E%3Cis_total_fee_adjust%3EN%3C%2Fis_total_fee_adjust%3E%3Ctotal_fee%3E0.01%3C%2Ftotal_fee%3E%3Cgmt_payment%3E2014-10-16+14%3A57%3A13%3C%2Fgmt_payment%3E%3Cseller_email%3Efirecloud168%40gmail.com%3C%2Fseller_email%3E%3Cprice%3E0.01%3C%2Fprice%3E%3Cbuyer_id%3E2088002000448273%3C%2Fbuyer_id%3E%3Cnotify_id%3E721c8ccb6a7bf5c2ba30f45c8a5750a53i%3C%2Fnotify_id%3E%3Cuse_coupon%3EN%3C%2Fuse_coupon%3E%3C%2Fnotify%3E"
	sSign	=alipayMd5Sign(sStr)
	log.Printf("Sign(730581744159c4e24c083c20b23bfc19):%s", sSign)
}