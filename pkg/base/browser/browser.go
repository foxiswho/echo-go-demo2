package browser

type IBrowser interface {
}

// Browser 浏览器
type Browser struct {
	/**
	 * 操作系统
	 */
	os string

	/**
	 * 浏览器类型
	 */
	browser string
}
