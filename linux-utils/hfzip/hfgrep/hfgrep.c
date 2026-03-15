#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
  char buf[256];
  FILE *fptr;

  //If no args are passed
  if(argc < 2){
    fputs("hfgrep: searchterm [file ...]\n", stdout);
    return 1;
  }
  //If only a search term is passed
  if(argc == 2) {
    while(fgets(buf, sizeof(buf), stdin)) {
      if(strstr(buf,argv[1])) {
        fputs(buf, stdout);
      }
    }
  }
  //If file(s) and a search term have been passed
  for(int i = 2; (argv[i]); i++) {
    fptr = fopen(argv[i], "r");
    if(!(fptr)) {
      fputs("hfgrep: cannot open file\n", stdout); 
      return 1;
    }
    while(fgets(buf, sizeof(buf), fptr)) {
      if(strstr(buf,argv[1])) {
        fputs(buf, stdout);
      }
    }
  fclose(fptr);
  }
}
