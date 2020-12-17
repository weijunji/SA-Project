<template>
  <v-app>
    <v-navigation-drawer
      class="deep-purple accent-4"
      app
      dark
      permanent
    >
      <v-list>
        <v-list-item>
          分布式监管系统
        </v-list-item>
        <v-divider/>
        <v-list-item
          v-for="item in items"
          :key="item.title"
          link
          :href="item.link"
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>
              {{ item.title }}
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      items: [
        { title: '首页', icon: 'mdi-home', link: '/home' },
        { title: '分布式机器', icon: 'mdi-server', link: '/clients' },
        // { title: '帮助', icon: 'mdi-help-circle' },
      ],
    }
  },
};
</script>
