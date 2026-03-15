#include <stdio.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
  FILE *fptr;
  int count;
  char c;

  //check when no args are passed
  if(argc < 2) {
    fputs("hfunzip: file1 [file2 ...]\n",stdout);
    return 1;
  }

  //loop for every file passed
  for(int i = 1; (argv[i]); i++) {
    fptr = fopen(argv[i], "r");
    if(!fptr) {
      fputs("Error opening file",stdout);
      return 1;
    }
    
    //reads every 4 bytes and stores into count 
    while(fread(&count, 1, sizeof(int),fptr) == sizeof(int)) {
      //reads last 1 byte and stores into c
      if(fread(&c, 1, sizeof(char), fptr) != sizeof(char)) {
        return 1;
      }
      //prints chars based on count read
      while(count > 0) {
        printf("%c",c);
        count--;
      }
    }
    fclose(fptr);
  }
}
