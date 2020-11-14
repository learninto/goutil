# xhttp


## http

```go
const Timeout = 6 * time.Second
const ApiFindUserById = Root + "/twirp/user.v1.User/FindUserById"

    params := struct {
		Id int64 `json:"id"`
	}{
		Id: id,
	}
	b, _ := json.Marshal(params)

	req, _ := http.NewRequest(http.MethodPost, frame.ApiFindUserById, bytes.NewReader(b))

	resp, err := httpD.NewClient(frame.Timeout).Do(ctx, req)
	fmt.Println(resp)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		return u, errors.New("请求失败")
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(respBody, &u)
```

## grpc

支持 grpc unary 请求

grpc 的应用日益广泛，但 grpc 又过于复杂。就我们观察，绝大多数 grpc 接口
都是 unary 类型的。此类请求跟传统的 http 请求很类似，比较简单。

基于此，我们引入一个简单的 unary client 实现，争取用 20% 的代码解决 80%
的问题。
