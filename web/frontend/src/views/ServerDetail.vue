<template>
  <div class="about" style="margin: 10px">
    <h1>客户端管理：</h1>
    <v-card style="margin-top: 10px">
        <v-card-title>
            {{name}}
            <v-btn icon small color="cyan" @click="dialog = true" style="margin-left: 5px">
                <v-icon>mdi-pencil</v-icon>
            </v-btn>
        </v-card-title>
        <v-card-subtitle>{{uuid}}</v-card-subtitle>
        <div>
            <v-btn @click="lockHandle" :color="locked ? 'green' : 'red'" style="margin: 10px">
                {{locked ? 'Locked' : 'Unlock'}}
            </v-btn>
            <v-btn @click="confirmDialog = true" :disabled="!online" :color="online ? 'blue' : 'grey darken-2'" style="margin: 10px">
                {{online ? 'Online' : 'Offline'}}
            </v-btn>
        </div>
    </v-card>
    <v-card style="margin-top: 20px">
        <v-card-title>
            操作日志
            <v-btn icon small color="cyan" @click="downloadLogs" style="margin-left: 5px">
                <v-icon>mdi-download</v-icon>
            </v-btn>
        </v-card-title>
        <v-card-text>
            <v-data-table
                :headers="headers"
                :items="logs"
                group-by="Operation"
                disable-pagination
                fixed-header
                height="400px"
            ></v-data-table>
        </v-card-text>
    </v-card>
    <v-snackbar
      v-model="snackbar"
    >
      {{ text }}
      <template v-slot:action="{ attrs }">
        <v-btn
          color="pink"
          text
          v-bind="attrs"
          @click="snackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
    <v-dialog v-model="dialog" persistent max-width="290">
      <v-card>
        <v-card-title class="headline">修改名称为：</v-card-title>
        <v-card-text>
            <v-text-field v-model="editName" label="Client Name"></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" text @click="dialog = false">取消</v-btn>
          <v-btn color="green darken-1" text @click="editNameHandle">确认</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="confirmDialog" persistent max-width="290">
      <v-card>
        <v-card-title class="headline">确认关闭该机器？</v-card-title>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" text @click="confirmDialog = false">取消</v-btn>
          <v-btn color="red darken-1" text @click="onlineHandle">确认</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
    name: 'ServerDetail',
    data () {
        return {
            uuid: "",
            name: "",
            locked: true,
            online: false,
            logs: [],
            snackbar: false,
            text: "操作成功",
            dialog: false,
            editName: "",
            confirmDialog: false,
            headers: [
                { text: '编号', value: 'ID' },
                { text: '时间', value: 'CreatedAt', sortable: false },
                { text: '操作', value: 'Operation' },
            ],
        }
    },
    methods: {
        lockHandle () {
            this.$axios.put(`http://127.0.0.1/client/${this.uuid}`, {
                locked: !this.locked
            }).then(response => {
                if(response.status == 200) {
                    this.locked = !this.locked
                    this.snackbar = true
                }
            })
        },
        onlineHandle () {
            if (this.online) {
                this.$axios.put(`http://127.0.0.1/client/${this.uuid}`, {
                    online: false
                }).then(response => {
                    if(response.status == 200) {
                        this.online = false
                        this.text = "操作成功"
                        this.snackbar = true
                    }
                })
            }
            this.confirmDialog = false
        },
        editNameHandle () {
            this.$axios.put(`http://127.0.0.1/client/${this.uuid}`, {
                name: this.editName
            }).then(response => {
                if(response.status == 200) {
                    this.name = this.editName
                    this.text = "操作成功"
                    this.snackbar = true
                }
                this.dialog = false
            })
        },
        download (filename, text) {
            var element = document.createElement('a');
            element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
            element.setAttribute('download', filename);

            element.style.display = 'none';
            document.body.appendChild(element);

            element.click();

            document.body.removeChild(element);
        },
        downloadLogs () {
            let f = "ID,Time,Operation\n"
            for (let log of this.logs) {
                f += `${log.ID},${log.CreatedAt},${log.Operation}\n`
            }
            this.download(`${this.uuid}.csv`, f)
        }
    },
    created () {
        this.uuid = this.$route.params.id
        this.$axios.get('http://127.0.0.1/client').then(response => {
            for(let client of response.data){
                if(client.UUID == this.uuid){
                    this.name = client.Name
                    this.locked = client.Locked
                    this.online = client.Online
                }
            }
        })
        this.$axios.get(`http://127.0.0.1/client/logs/${this.uuid}`).then(response => {
            this.logs = response.data
            const m = ["Lock", "Unlock", "Online", "Offline"]
            for(let i in this.logs) {
                this.logs[i].ID = Number(i) + 1
                this.logs[i].CreatedAt = (new Date(Date.parse(this.logs[i].CreatedAt))).toLocaleString()
                let op = this.logs[i].Operation
                this.logs[i].Operation = m[op]
            }
            console.log(this.logs)
        })
    }
}
</script>
