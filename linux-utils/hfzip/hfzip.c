#include <stdio.h>
#include <unistd.h>

//writes binar+char sequence to stdout
void writeSeq(int *n, char *c){
  fwrite(n, sizeof(int), 1, stdout);
  fwrite(c, sizeof(char), 1, stdout);
  *n = 1;
}

int main(int argc, char *argv[]) {
  char buf;
  char mem;
  FILE *fptr;
  int count = 0;

  //checks when no args are passed
  if(argc < 2) {
    fputs("hfzip: file1 [file2 ...]\n",stdout);
    return 1;
  }

  //reads through every argument and places into one zip
  for(int i = 1; (argv[i]); i++) {
    fptr = fopen(argv[i], "r");
    if(!fptr) {
      fputs("hfzip: cannot open file\n", stdout);
      return 1;
    }
    //memory for comparison
    mem = fgetc(fptr);
    if(mem != EOF) {
      count++;
    }
    //reads each char in file
    while((buf = fgetc(fptr)) != EOF ) {
      if(mem == buf) {
        count++;
      } else {
        writeSeq(&count, &mem);
      }
      mem = buf;
    }
    //checks when there are more files
    if(count > 0 && mem != buf && i == (argc-1)) {
      writeSeq(&count, &mem);
    }
    fclose(fptr);
  }
  return 0;
}
