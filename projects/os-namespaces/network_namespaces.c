#define _GNU_SOURCE
#define STACK_SIZE (1024 * 1024)

#include <sched.h>
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <sys/socket.h>
#include <sys/wait.h>

static char child_stack[STACK_SIZE];

static int child_fn(void *arg) {
  printf("Namespace:\n");
  system("ip link");
  printf("\n\n");
  sleep(100);
  return 0;
}

void make_new_network_namespace() {
  pid_t child_pid = clone(child_fn, child_stack+STACK_SIZE, CLONE_NEWNET | SIGCHLD, NULL);
  if (child_pid != 0)
    perror("clone command failed");
  waitpid(child_pid, NULL, 0);
}

void dns_server() {
  int server_fd = socket(AF_INET, SOCK_DGRAM, 0);
  struct sockaddr_in serv_addr;
  if (server_fd == -1)
    printf("Could not open server socket");

}

// change the current namespace to an existing namespace and connecting to a
// virtual interface from that namespace
void change_namespace() {
  int new_ns;
}

int main() {
  /* make_new_network_namespace(); */
  change_namespace();
}
