extern "C" {
    double heron(const double number) {
        double x = 42.0;
        for (int i = 0; i < 1000; ++i) {
            x = (x + number / x) / 2.0;
        }
        return x;
    }

    double heron_int(const int number) {
        return heron(number);
    }
}