import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";
import {useUserStore} from "../store/user.ts";

const routes : RouteRecordRaw[] = [
    {
        path: "/login",
        component: () => import("../views/login.vue"),
    },
    {
        path: "/main/:pathMatch(.*)*",
        component: () => import("../views/main.vue"),
        alias: "/",
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach((to, _, next) => {
    const userStore = useUserStore()
    if (to.path !== "/login" && to.path !== '/auth' && !userStore.isLogin) {
        sessionStorage.setItem("redirect", to.path)
        next("/login")
    } else {
        next()
    }
})

export default router