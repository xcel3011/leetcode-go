package leetcode_go

//一个括号字符串是只由 '(' 和 ')' 组成的 非空 字符串。如果一个字符串满足下面 任意 一个条件，那么它就是有效的：
//
//字符串为 ().
//它可以表示为 AB（A 与 B 连接），其中A 和 B 都是有效括号字符串。
//它可以表示为 (A) ，其中 A 是一个有效括号字符串。
//给你一个括号字符串 s 和一个字符串 locked ，两者长度都为 n 。locked 是一个二进制字符串，只包含 '0' 和 '1' 。对于 locked 中 每一个 下标 i ：
//
//如果 locked[i] 是 '1' ，你 不能 改变 s[i] 。
//如果 locked[i] 是 '0' ，你 可以 将 s[i] 变为 '(' 或者 ')' 。
//如果你可以将 s 变为有效括号字符串，请你返回 true ，否则返回 false 。

func canBeValid(s string, locked string) bool {
}
