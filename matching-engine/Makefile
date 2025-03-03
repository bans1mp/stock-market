CXX = g++
CXXFLAGS = -std=c++17 -I/usr/local/include -Iproto -Isrc
LDFLAGS = -L/usr/local/lib -lgrpc++ -lgrpc -lprotobuf -lgpr -lhiredis -labsl_synchronization -labsl_time -labsl_strings -labsl_base -pthread

SRC = src/server.cpp proto/trade.pb.cc proto/trade.grpc.pb.cc src/order_book.cpp src/redis.cpp src/main.cpp
OUT = build/matching_engine

all: $(OUT)

$(OUT): $(SRC)
	mkdir -p build
	$(CXX) $(CXXFLAGS) -o $(OUT) $(SRC) $(LDFLAGS)

clean:
	rm -rf build/
