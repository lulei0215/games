## GVA 
#### ：GIN-VUE-ADMIN 

### 

#### 1. GVAinitialize/router.go Routers 
    ：
    gva  gva 
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
		))

    

    PluginInit(PrivateGroup, email.CreateEmailPlug(
    "a@qq.com",
    "b@qq.com",
    "smtp.qq.com",
    "global.GVA_CONFIG.Email.Secret",
    "",
    465,
    true,
    ))

### 2. 

#### 2-1 
    // Form  Secret 

    type Email struct {
	    To       string  // : ：a@qq.com b@qq.com  
	    From     string  //   
	    Host     string  //   smtp.qq.com  QQsmtp
	    Secret   string  //       smtp
	    Nickname string  //       
	    Port     int     //      QQsmtp  465
	    IsSSL    bool    // SSL   SSL
    }
#### 2-2 
    // Form  Secret 

    type Email struct {
        To      string `json:"to"`      // 
        Subject string `json:"subject"` // 
        Body    string `json:"body"`    // 
    }


### 3. API

    utils.EmailTest(，) 
    :utils.EmailTest(""，"")
    utils.ErrorToEmail(,) 
    :utils.ErrorToEmail(""，"")
    utils.Email(，，) 
    :utils.Email(”a.qq.com,b.qq.com“,""，"")

### 4. 

    ： /email/emailTest [post] swagger

    ： /email/emailSend [post] swagger
    ：
    type Email struct {
        To      string `json:"to"`      // 
        Subject string `json:"subject"` // 
        Body    string `json:"body"`    // 
    }
   
