<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>, Félix PORTIER <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <Modal id="admin-newpage" :title="$t('sidebar__create_new_page')" customClass="modal fade in">
    <form @submit.prevent="savePage" @keydown="form.onKeydown($event)">

      <!-- Name -->
      <div class="form-group row">
        <label class="col-md-4 col-form-label text-md-right">
          {{ $t('settings_pages__name') }}
        </label>
        <span class="col-md-8">
          <input v-model="form.name" type="text" name="name" class="form-control">
        </span>
      </div>

      <hr>

     <!-- Optionnal -->
      <div class="form-group row">
        <label class="col-md-4 col-form-label">
          {{ $t('create_page__optional') }}
        </label>
      </div>

      <!-- Title -->
      <div class="form-group row">
        <label class="col-md-4 col-form-label text-md-right">
          {{ $t('settings_pages__title') }}
        </label>
        <span class="col-md-8">
          <input v-model="form.title" type="text" name="title" class="form-control">
        </span>
      </div>

      <!-- Keywords -->
      <div class="form-group row">
        <label class="col-md-4 col-form-label text-md-right">
          {{ $t('settings_pages__keywords') }}
        </label>
        <span class="col-md-8">
          <input v-model="form.keywords" type="text" name="keywords" class="form-control">
        </span>
      </div>

      <!-- Description -->
      <div class="form-group row">
        <label class="col-md-4 col-form-label text-md-right">
          {{ $t('settings_pages__description') }}
        </label>
        <span class="col-md-8">
          <textarea v-model="form.description" name="description" class="form-control"></textarea>
        </span>
      </div>
    </form>

    <!-- Create button -->
    <div class="form-group row">
      <span class="col-md-4 offset-md-4 d-flex justify-content-start">
        <v-button :loading="form.busy" slot="footer" @click="savePage" class="btn btn-success">
          {{ $t('button__create') }}
        </v-button>
        </span>
      </div>

  </Modal>
</template>

<script>
import { mapGetters } from 'vuex'
import Form from 'vform'
import $ from 'jquery'
export default {
  name: 'admin-newpage',

  mounted () {
    this.$nextTick(() => {
      if (this.$route.query.new) {
        this.launchModal()
      }
    })
  },

  data () {
    return {
      form: new Form({
        name: '',
        title: '',
        link: '',
        keywords: '',
        description: ''
      })
    }
  },

  computed: mapGetters({
  }),

  beforeDestroy () {
    // destroy modal on component destroy
    this.destroyModal()
  },

  methods: {
    // Set form into default value
    initForm () {
      this.form = new Form({
        name: '',
        title: '',
        link: '',
        keywords: '',
        description: ''
      })
    },
    // Save new page
    savePage () {
      this.$validator.validateAll()
        .then(async (res) => {
          if (!this.form.name || this.form.name.trim() === '') {
            this.$awn.alert(this.$t('settings_pages__empty_page_name'))
          } else {
            this.form.name = this.form.name.trim()
            if (res) {
              try {
                const { data } = await this.form.post(process.env.API_URL + '/api/page')
                this.initForm()
                this.fetchAllPage()
                this.$awn.success(this.$t('create_page__page_created_1') + data.name + this.$t('create_page__page_created_2'))
                this.closeModal()
                this.$router.replace('/' + data.link)
              } catch (e) {
                this.$awn.alert(this.$t(e.response.data.message))
              }
            }
          }
        })
    },
    async fetchAllPage () {
      await this.$store.dispatch('page/fetchAllPage')
    },
    // launch modal
    launchModal () {
      $('#admin-newpage').modal('show')
    },
    // close modal
    closeModal () {
      $('#admin-newpage').modal('hide')
    },
    // destroy modal
    destroyModal () {
      $('#admin-newpage').modal('dispose')
    }
  }
}
</script>
