go build clock.go
go build clockwall.go
TZ=America/New_York ./clock -port 8010 &
TZ=Asia/Tokyo ./clock -port 8020 &
TZ=Europe/London ./clock -port 8030 &
./clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030&
