<template>
  <div>
    <h1>View Product {{ product._id }}</h1>
    <p>{{ product.title }}</p>
    <p>Price: {{ product.price }}</p>
    <div id="app">
      <el-button @click="visible = true">Button</el-button>
      <el-dialog :visible.sync="visible" title="Hello world">
        <p>Try Element</p>
      </el-dialog>
    </div>
  </div>

</template>

<script>
import { mapState } from 'vuex'
export default {
  created () {
    this.$store.dispatch('product/load')
  },
  computed: {
    ...mapState(['product'])
  }
}
</script>