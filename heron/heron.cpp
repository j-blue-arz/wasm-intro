extern "C" {
    double heron(const int number) {
        double x = 42.0;
        for (int i = 0; i < 1000; ++i) {
            x = (x + number / x) / 2.0;
        }
        return x;
    }
}