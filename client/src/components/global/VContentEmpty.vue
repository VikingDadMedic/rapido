<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <div class="row">
    <div class="col-md-12 py-3">
      <Floataddcontent :index="0" :location="location" @add="addNewContent" @addExisting="addExistingContent" @addExtension="addExtension" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'VContentEmpty',

  props: {
    location: {
      type: String,
      required: true
    },
    index: {
      type: Number,
      required: true
    }
  },

  methods: {
    addNewContent (v) {
      this.$emit('add', v, this.index)
    },
    addExistingContent (v) {
      this.$emit('addExisting', v, this.index)
    },
    addExtension (v) {
      this.$emit('addExtension', v, this.index)
    }
  }
}
</script>

<style scoped>
/* .row { margin-top: -15px; } */
/* .row { transform: translateY(-50%); } */
</style>
