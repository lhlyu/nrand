// 懒加载

const routers = [
  {
    path:"/",
    meta:{title:"首页"},
    component:(resolve) => require(["../components/Index.vue"], resolve)
  },
]

export default new VueRouter({
  routes:routers
})






