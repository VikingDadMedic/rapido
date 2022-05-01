<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <Modal id="gallery-modal" :title="$t('settings_files')" modalSize="modal-lg" customClass="modal fade in">
    <Gallery :asModal="true" />
  </Modal>
</template>

<script>
import { mapGetters } from 'vuex'
import $ from 'jquery'
import Gallery from '@/components/admin/Settings/Files.vue'
export default {
  name: 'gallery-modal',

  components: {
    Gallery
  },

  mounted () {
    this.$nextTick(() => {
      if (this.$route.query.new) {
        this.launchModal()
      }
    })
  },

  data () {
    return {
    }
  },

  computed: mapGetters({
    access_level: 'page/access_level'
  }),

  watch: {
    '$el.classList.value' () {
      // console.log('berubah')
    }
  },

  beforeDestroy () {
    // destroy modal on component destroy
    this.destroyModal()
  },

  methods: {
    // launch modal
    launchModal () {
      $('#gallery-modal').modal('show')
    },
    // close modal
    closeModal () {
      $('#gallery-modal').modal('hide')
    },
    // destroy modal
    destroyModal () {
      $('#gallery-modal').modal('dispose')
    }
  }
}
</script>
