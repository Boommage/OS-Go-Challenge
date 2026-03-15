#include <stdio.h>

int main(int argc, char *argv[]) {
  char buf[256];
  FILE *fptr;
  for(int i = 1; (argv[i]); i++)
  {
    fptr = fopen(argv[i],"r");
    if(!(fptr)) {
      //snprintf(buf,sizeof(buf), "cat: %s: No such file or directory\n", argv[i]);
      //fputs(buf, stdout);
      fputs("hfcat: cannot open file\n", stdout);
      return 1;
    }
    while(fgets(buf, sizeof(buf), fptr)) {
      //buf[strcspn(buf, "\r\n")] = '\0';
      //printf("%s\n",buf);
      fputs(buf, stdout);
    }
    fclose(fptr);
  }

  return 0;
}
