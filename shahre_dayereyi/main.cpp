#include <iostream>
using namespace std;

int main() {
    int n, m;
    cin >> n >> m;

    bool hasInward = false;
    bool hasOutward = false;

    for (int i = 0; i < n; i++) {
        int direction;
        cin >> direction;

        if (direction == 0) {
            hasInward = true;
        } else {
            hasOutward = true;
        }
    }

    // جهت خیابان‌های دایره‌ای تأثیری در جواب ندارد.
    for (int i = 0; i < m; i++) {
        int direction;
        cin >> direction;
    }

    if (hasInward && hasOutward) {
        cout << "YES" << endl;
    } else {
        cout << "NO" << endl;
    }

    return 0;
}