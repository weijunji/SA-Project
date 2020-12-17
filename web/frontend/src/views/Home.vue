<template>
  <div class="home text-center" style="margin-top: 200px">
    <div class="text-h3">欢迎使用分布式监管系统</div>
    <div class="text-h5 font-weight-bold" style="margin-top: 20px">作者：</div>
    <div v-for="author in authors" :key="author.id" class="text-h6">
      <span>{{author.name}}</span>
      <span> {{author.id}}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data () {
    return {
      authors: [
        {"name": "魏俊吉", "id": "SA20225474"},
        {"name": "朱浩", "id": "SA20225646"},
        {"name": "叶芳彪", "id": "SA20225564"},
        {"name": "姜涵", "id": "SA20225218"},
      ]
    }
  }
}
</script>
