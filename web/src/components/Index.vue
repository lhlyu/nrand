<template>
  <div>
    <el-row>
      <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <div class="grid-content bg-purple">
          <el-select v-model="req.target" placeholder="选择生成类型" style="width: 100%">
            <el-option-group
              v-for="group in options"
              :key="group.label"
              :label="group.label">
              <el-option
                v-for="item in group.options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-option-group>
          </el-select>
        </div>
      </el-col>
    </el-row>
    <el-row class="top10px">
      <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <div class="grid-content bg-purple">
          <el-input
            placeholder="自定义开头"
            v-model="req.prefix"
            clearable>
          </el-input>
        </div>
      </el-col>
    </el-row>
    <el-row class="top10px">
      <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <div class="grid-content bg-purple">
          <el-input
            placeholder="自定义结尾"
            v-model="req.suffix"
            clearable>
          </el-input>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <div class="grid-content bg-purple">
          <el-divider content-position="left">生成数量</el-divider>
          <el-input-number v-model="req.number" :min="1" :max="100" label="生成数量"  style="width: 100%"></el-input-number>
        </div>
      </el-col>
    </el-row>
    <el-row v-if="display">
      <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <div class="grid-content bg-purple">
          <el-divider content-position="left">名字长度</el-divider>
          <el-input-number v-model="req.length" :min="1" :max="5" label="名字长度"  style="width: 100%"></el-input-number>
        </div>
      </el-col>
    </el-row>

    <el-row class="top10px">
      <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <div class="grid-content bg-purple">
          <el-button @click="getTags" style="width: 100%">开始生成</el-button>
        </div>
      </el-col>
    </el-row>

    <el-divider></el-divider>
    <el-row>
      <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
        <div class="grid-content bg-purple">
          <el-tag v-for="v,i in tags" v-bind:key="i" :type="getTagType()" v-clipboard:copy="v">{{v}}</el-tag>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>



export default {
  name: 'Index',
  data () {
    return {
      options: [{
        label: '天朝单姓',
        options: [{
          value: 10000,
          label: '天朝单姓(男)'
        }, {
          value: 10001,
          label: '天朝单姓(女)'
        }]
      },{
        label: '天朝复姓',
        options: [{
          value: 10002,
          label: '天朝复姓(男)'
        }, {
          value: 10003,
          label: '天朝复姓(女)'
        }]
      },{
        label: '日本姓氏',
        options: [{
          value: 10004,
          label: '日本姓氏(男)'
        }, {
          value: 10005,
          label: '日本姓氏(女)'
        }]
      },{
        label: '西方姓氏',
        options: [{
          value: 10006,
          label: '西方姓氏(男)'
        }, {
          value: 10007,
          label: '西方姓氏(女)'
        }]
      },{
        label: '其他',
        options: [{
          value: 10008,
          label: '组织'
        },{
          value: 10009,
          label: '地名'
        }, {
          value: 10010,
          label: '招式'
        }, {
          value: 10011,
          label: '天材地宝'
        }]
      }],
      tags:[],
      typs:["","success","info","warning","danger"],
      filters: [10006,10007,10010,10011],
      req: {
        target: 10000,
        prefix: "",
        suffix: "",
        number: 10,
        length: 1
      }
    }
  },
  methods:{
    getTagType(){
      let r = Math.random()*5
      let index = Math.round(r)
      return this.typs[index]
    },
    getTags(){
      let that = this
      this.tags = []
      let url = "/api/name"
      this.axios.post(url,this.req).then(function (response) {
        if(response.status == 200){
          let result = response.data
          if(result && result.data && result.data.length){
            that.tags = result.data
          }else{
            console.log(result)
          }
        }else{
          alert("工厂爆炸了，生产不了了！！！")
        }
      }).catch(function (error) {
        alert("你不小心捅破了服务器！！！")
        console.log(error);
      });
    }
  },
  computed:{
    display(){
      return this.filters.indexOf(this.req.target) == -1
    },
  }
}
</script>

<style scoped>
  .top10px{
    margin-top: 10px;
  }
  .el-tag{
    margin-right: 5px;
    margin-bottom: 5px;
  }
</style>
