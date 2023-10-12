import { defineStore } from "pinia"
import type {User} from "../types/user";

export const useUserStore = defineStore({
    id: "user",
    state: () : {
        user: User,
        token: string,
    } => ({
        user: {
            nickName: "",
        },
        token: "",
    }),
    getters: {
        isLogin: (state) => {
            return state.token !== ""
        }
    },
    actions: {
        login({token, user}: {token: string, user: User}) {
            this.token = token
            this.user = user
        }
    },
    persist: {
        paths: ['token']
    }
})

