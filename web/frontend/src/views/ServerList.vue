<template>
  <div class="server-list">
    <h1 style="margin-left: 10px">分布式机器列表</h1>
    <div v-for="server in servers" :key="server.UUID" style="margin: 10px 10px">
        <ServerCard v-if="server.UUID != 'web-server'" :uuid="server.UUID" :name="server.Name" :locked="server.Locked" :online="server.Online"></ServerCard>
    </div>
  </div>
</template>

<script>
import ServerCard from "@/components/ServerCard.vue"

export default {
    name: 'ServerList',
    components: {ServerCard},
    data () {
        return {
            servers: []
        }
    },
    created () {
        this.$axios.get("http://127.0.0.1/client").then(response => {
            this.servers = response.data
        })
    }
}
</script>
