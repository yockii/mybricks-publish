<script setup lang="ts">
import {onMounted, ref} from "vue";
import {loadMicroApp, MicroApp} from "qiankun";
import {User} from "../types/user";
import {useUserStore} from "../store/user";
import router from "../router";
const userStore = useUserStore()
const loginContainer = ref<HTMLElement | null>(null)

const logined = ({token, user}: {token: string, user: User}) => {
  userStore.token = token
  userStore.user = user
  router.push('/main')
}

const loginApp = ref<MicroApp | null>(null)
onMounted(() => {
  if (userStore.isLogin) {
    router.push('/main')
    return
  }


  if (!window.logined) {
    window.logined = logined
  }
  if (loginContainer.value) {
    loginApp.value = loadMicroApp({
      name: 'login',
      entry: 'asset/12877708147556352/index.html',
      container: loginContainer.value,
    }
    )
  }
})
</script>

<template>
  <div ref="loginContainer" class="w-full h-full"></div>
</template>

<style>
#__qiankun_microapp_wrapper_for_login__ {
  height: 100%;
}
</style>