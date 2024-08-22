<script setup>
import {onMounted} from "vue";

onMounted(() => {
  var bd = document.getElementById("my-body")
  bd.style.width = window.innerWidth + 'px'
  bd.style.height = window.innerHeight+'px'
})
</script>

<template>
  <div id="my-body">
<!--    123123-->
    <router-view/>
  </div>
</template>

<style scoped>

</style>