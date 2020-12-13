vector<int> partitionLabels(string S) {
    vector<int> result;
    unordered_map<char, int> map; //记录char c 和其最后出现位置的 map
    int start = 0, end = 0;
    for (int i = 0; i < S.size(); i ++) {
        map[S[i]] = i;
    }
    for (int i = 0; i < S.size(); i ++) {
        end = max(end, map[S[i]]);
        if (i == end) {
            result.push_back(end - start + 1);
            start = i + 1;
        }
    }
    return result;
}