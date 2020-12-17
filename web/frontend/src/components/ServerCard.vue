<template>
    <v-card :href="'/clients/'+uuid">
        <v-list-item tow-line>
            <v-list-item-content>
                <v-list-item-title>{{name}}</v-list-item-title>
                <v-list-item-subtitle>{{uuid}}</v-list-item-subtitle>
            </v-list-item-content>
            <v-list-item-icon>
                <v-chip :color="locked ? 'green' : 'red'" text-color="white">
                    <v-avatar left>
                        <v-icon small>{{locked ? 'mdi-lock-outline' : 'mdi-lock-open-outline'}}</v-icon>
                    </v-avatar>
                    {{locked ? 'Locked' : 'Unlock'}}
                </v-chip>
            </v-list-item-icon>
            <v-list-item-icon>
                <v-chip :color="online ? 'blue' : 'grey darken-2'" text-color="white">{{online ? 'Online' : 'Offline'}}</v-chip>
            </v-list-item-icon>
        </v-list-item>
    </v-card>
</template>

<script>
export default {
    name: "ServerCard",
    props: [
        "uuid", "name", "online", "locked"
    ]
}
</script>
