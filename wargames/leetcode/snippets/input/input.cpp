#include <stdio.h>
#include <string.h>

#include <iostream>
#include <vector>
using namespace std;

void read_n_lines() {
  int n;
  string s;
  vector<string> lines;
  cin >> n;
  cout << "n: " << n << endl;
  for (int i = 0; i < n; ++i) {
    cin >> s;
    lines.push_back(s);
  }
  for (auto s : lines) {
    cout << s << endl;
  }
}

void read_until_newline() {
  string s;
  vector<string> lines;
  while (1) {
    getline(cin, s);
    if (s.length() == 0) break;
    lines.push_back(s);
  }
  for (auto s : lines) {
    cout << s << endl;
  }
}

int main() {
  int n;
  read_until_newline();
  return 0;
}
