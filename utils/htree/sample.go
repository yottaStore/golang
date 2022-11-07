package htree

import "log"

func SampleTree() (*Root, error) {

	r, err := New("root.com")
	if err != nil {
		log.Fatal("Error creating tree: ", err)
	}

	us := []Update{
		{
			Type:   INSERT_SHARD,
			Url:    "s2.n1.dc1.root.com:8081",
			Weight: 5,
			Load:   0,
		},
		{
			Type:   INSERT_SHARD,
			Url:    "s1.n1.dc1.root.com:8081",
			Weight: 10,
			Load:   0,
		},
		{
			Type:   INSERT_SHARD,
			Url:    "s3.n1.dc1.root.com:8081",
			Weight: 5,
			Load:   0,
		},
		{
			Type:   INSERT_SHARD,
			Url:    "s1.n2.dc1.root.com:8081",
			Weight: 10,
			Load:   0,
		},
		{
			Type:   INSERT_SHARD,
			Url:    "s2.n2.dc1.root.com:8081",
			Weight: 10,
			Load:   0,
		},
		{
			Type:   UPDATE_LOAD,
			Url:    "n1.dc1.root.com",
			Weight: 0,
			Load:   5,
		},
		{
			Type:   UPDATE_LOAD,
			Url:    "n2.dc1.root.com",
			Weight: 0,
			Load:   5,
		},
	}

	err = r.Update(us)

	//r.Print()

	if err != nil {
		return nil, err
	}

	err = r.Verify()
	if err != nil {
		return nil, err
	}

	return r, nil
}
