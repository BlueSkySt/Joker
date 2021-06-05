    if err != nil || cmd == "PUZ1" || cmd == "puz1" {
    this.conn.Write([]byte("\033[2J\033[1H")) //display main header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;37m  \033[01;31mnice choice! \033[01;37m" + username + " !          \r\n")))
            this.conn.Write([]byte("\033[37m                      \033[01;31m_           \r\n"))
            this.conn.Write([]byte("\033[37m\033[37m  _ __ \033[01;31m _   _ \033[37m____\033[01;31m___\033[37m| |\033[01;31m ___ \033[37m ___ \r\n"))
            this.conn.Write([]byte("\033[37m\033[37m | '_ \\033[01;31m| | | \033[37m|_  /\033[01;31m_  \033[37m/ |\033[01;31m/ _ \\033[37m/ __|\r\n"))
            this.conn.Write([]byte("\033[37m\033[37m | |_) \033[01;31m| |_| |\033[37m/ / \033[01;31m/ /\033[37m| |\033[01;31m  __/\033[37m|__ \\r\n"))
            this.conn.Write([]byte("\033[37m\033[37m | .__/ \033[01;31m|__,_/\033[37m___/\033[01;31m___\033[37m|_|\033[01;31m|___|\033[37m|___/\r\n"))
            this.conn.Write([]byte("\033[37m\033[37m |_|                            \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
  continue
    } 
