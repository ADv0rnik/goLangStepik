// вы уже внутри main()
done := make(chan struct{})
go func(c chan struct{}) {
	work()
	close(c)
}(done)
<-done
