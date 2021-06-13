#include <iostream>
#include <map>
#include <sstream>
#include <vector>

using namespace std;

typedef struct {
  int n;
  vector<vector<int>> edges;
} InputData;

InputData read_data() {
  vector<vector<int>> data;
  string line;
  int n, i = 0;
  cin >> n;
  getline(cin, line);
  while (1) {
    int a, b;
    stringstream ss;
    getline(cin, line);
    if (line.empty()) break;
    ss << line;
    ss >> a >> b;
    data.push_back({a, b});
  }

  /* for (auto e : data) { */
  /*   for (auto n : e) { */
  /*     cout << n << " "; */
  /*   } */
  /*   cout << endl; */
  /* } */
  return InputData{n, data};
}

int findJudgeMap(int n, vector<vector<int>>& trust) {
  map<int, int> m;
  for (auto i : trust) {
    m[i[1]]++;
    m[i[0]]--;
  }
  for (auto it : m) {
    if (it.second == n - 1) return it.first;
  }
  if (n == 1 && m.empty()) return 1;
  return -1;
}

int findJudge(int n, vector<vector<int>>& trust) {
  int i;
  int trustCount[1001] = {};
  for (auto i : trust) {
    trustCount[i[1]]++;
    trustCount[i[0]]--;
  }
  for (int i = 1; i <= n; i++) {
    if (trustCount[i] == n - 1) return i;
  }
  return -1;
}

int main() {
  cout << "--------------------------------" << endl;
  InputData data = read_data();
  cout << "Judge:" << findJudgeMap(data.n, data.edges) << endl;
}
