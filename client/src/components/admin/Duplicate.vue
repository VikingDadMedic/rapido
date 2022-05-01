<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>, Félix PORTIER <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <Modal id="admin-duplicate" :title="$t('sidebar__duplicate_page')" customClass="modal fade in">
    <form @submit.prevent="copyPage" @keydown="form.onKeydown($event)">

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

      <!-- Link -->
      <div class="form-group row">
        <label class="col-md-4 col-form-label text-md-right">
          {{ $t('settings_pages__link') }}
        </label>
        <span class="col-md-8">
          <input v-model="form.link" type="text" name="link" class="form-control" readonly>
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

    <!-- Duplicate button -->
    <div class="form-group row">
      <span class="col-md-4 offset-md-4 d-flex justify-content-start">
        <v-button :loading="form.busy" slot="footer" @click="copyPage" class="btn btn-success">
          {{ $t('button__duplicate') }}
        </v-button>
        </span>
      </div>

  </Modal>
</template>

<script>
import { mapGetters } from 'vuex'
import Form from 'vform'
import _ from 'lodash'
import $ from 'jquery'
export default {
  name: 'admin-duplicate',

  mounted () {
    this.$nextTick(() => {
      if (this.$route.query.duplicate) {
        this.launchModal()
        _.assign(this.form, this.pages.details)
      }
    })
  },

  data () {
    return {
      form: new Form({
        access_level: 1,
        name: '',
        title: '',
        link: '',
        keywords: '',
        description: ''
      })
    }
  },

  computed: mapGetters({
    pages: 'page/current_pages',
    access_level: 'page/access_level'
  }),

  watch: {
    pages: {
      handler (v) {
        _.assign(this.form, v.details)
      },
      deep: true
    }
  },

  beforeDestroy () {
    // destroy modal on component destroy
    this.destroyModal()
  },

  methods: {
    // Set form into default value
    initForm () {
      this.form = new Form({
        access_level: 1,
        name: '',
        title: '',
        link: '',
        keywords: '',
        description: ''
      })
    },
    // Save new page
    async copyPage () {
      this.$validator.validateAll()
        .then(async (res) => {
          if (!this.form.name || this.form.name.trim() === '') {
            this.$awn.alert(this.$t('settings_pages__empty_page_name'))
          } else {
            this.form.name = this.form.name.trim()
            if (res) {
              try {
                const { data } = await this.form.post(process.env.API_URL + '/api/page/' + this.pages.details.id + '/copy')
                this.initForm()
                this.fetchAllPage()
                this.$awn.success(this.$t('duplicate_page__page_duplicated_1') + data.name + this.$t('duplicate_page__page_duplicated_2'))
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
      $('#admin-duplicate').modal('show')
    },
    // close modal
    closeModal () {
      $('#admin-duplicate').modal('hide')
    },
    // destroy modal
    destroyModal () {
      $('#admin-duplicate').modal('dispose')
    }
  }
}
</script>
