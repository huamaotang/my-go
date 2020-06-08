## select的多路复用
### 一、知识点
    1、类似于switch语句，但select不会有输入值，仅用于信道操作
    2、作用：从多个发送或接受信道操作中进行选择，select语句会阻塞直到其中有信道可以操作
    如果有多个信道可操作，会随机选择其中一个case执行
    3、default case：使得select语句不再阻塞，如果其它信道操作还没准备好，将会执行default
    分支
    4、nil channel：信道的默认值是nil，不能对nil的信道进行读写操作；case分支中如果信道是
    nil，该分支就会被忽略，可以使用default case避免报错
    5、添加超时时间：可以在case语句后面设置超时时间；如果在设置的时间范围之内还没有可操作的信道
    则执行该分支
    6、空select：没有case语句的select，将会一直阻塞，发生死锁
    7、select机制用来处理异步IO问题
    8、select机制里最大的限制就是case中必须是一个IO操作
    9、golang在语言级别支持select关键字
