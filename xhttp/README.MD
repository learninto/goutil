
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