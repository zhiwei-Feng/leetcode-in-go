package n2029

func stoneGameIX(stones []int) bool {
	cnt := [3]int{}
	for _, v := range stones {
		cnt[v%3]++
	}
	return play(true, 0, cnt[0], cnt[1], cnt[2])
}

// params:
// removed(int): 表示当前删除的价值 mod 3的结果
// return: true->Alice win ; false->Bob win
func play(who bool, removed int, c0, c1, c2 int) bool {
	if c0 == 0 && c1 == 0 && c2 == 0 {
		// 没石子=>bob win
		return false
	}
	flag := !who // 假设当前角色who赢不了
	if removed != 0 {
		// 当游戏进行到第二回合及以后，无论是谁的回合，他的策略应该如下
		// 1. 有价值mod 3 == 0的石子就删除
		// 1不满足，则考虑当前删除的removed价值mod 3 是1还是2
		// 如果是1，则自己删除mod 3后为2的石子必输，如果是2同理
		// 因此，removed==1则删除c1, removed==2则删除c2
		if c0 > 0 {
			return play(!who, removed, c0-1, c1, c2)
		} else if removed == 1 {
			if c1 > 0 && play(!who, 2, c0, c1-1, c2) == who {
				flag = who
			}
		} else {
			if c2 > 0 && play(!who, 1, c0, c1, c2-1) == who {
				flag = who
			}
		}
	} else {
		// 刚开始的时候，alice可以删除c1或者c2, 穷举即可
		if c1 > 0 && play(!who, 1, c0, c1-1, c2) == who {
			flag = who
		}
		if c2 > 0 && play(!who, 2, c0, c1, c2-1) == who {
			flag = who
		}
	}
	return flag
}
