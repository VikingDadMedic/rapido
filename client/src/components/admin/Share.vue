<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <Modal id="admin-share" :title="$t('sidebar__share_page')" customClass="modal fade in">
    <form @submit.prevent="addAuthor" @keydown="form.onKeydown($event)">

      <!-- Your email address -->
      <div class="form-group row" v-if="!user">
        <label class="col-md-4 col-form-label text-md-right">
          {{ $t('your_email_address') }}
        </label>
        <span class="col-md-8">
          <input v-model="form.your_email" type="text" name="your_email" class="form-control">
        </span>
      </div>

      <!-- Recipient email address -->
      <div class="form-group row">
        <label class="col-md-4 col-form-label text-md-right">
          {{ $t('recipient_email_address') }}
        </label>
        <span class="col-md-8">
          <input v-model="form.email" type="text" name="email" class="form-control">
        </span>
      </div>

    </form>

      <!-- Share button -->
      <div class="form-group row">
        <span class="col-md-4 offset-md-4 d-flex justify-content-start">
          <v-button :loading="form.busy" slot="footer" @click="addAuthor" class="btn btn-success">
            {{ $t('button__share') }}
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
  name: 'admin-share',

  mounted () {
    this.fetchEmailSender()
    this.$nextTick(() => {
      if (this.$route.query.share) {
        this.launchModal()
      }
    })
  },

  data () {
    return {
      email_sender: String,
      regexpPattern: /^.+@.+$/,
      form: new Form({
        your_email: '',
        email: ''
      })
    }
  },

  computed: mapGetters({
    user: 'auth/user',
    pages: 'page/current_pages'
  }),

  beforeDestroy () {
    // destroy modal on component destroy
    this.destroyModal()
  },

  methods: {

    // fetch config from DB to test for email_sender presence
    async fetchEmailSender () {
      try {
        const data = await this.$store.dispatch('auth/fetchSetting')
        this.email_sender = data.email_sender
      } catch (e) {}
    },

    // Set form into default value
    initForm () {
      this.form = new Form({
        email: ''
      })
    },

    // Check if email addresses are valid using regexp, before sharing page
    async addAuthor () {
      if (!this.email_sender || !this.regexpPattern.test(this.email_sender.trim())) {
        this.$awn.alert(this.$t('settings_system__invalid_sender_email_address'))
      // User isn't connected and had to enter his email address
      } else if (!this.user && (!this.form.your_email || !this.regexpPattern.test(this.form.your_email.trim()))) {
        this.$awn.alert(this.$t('your_email_address_is_invalid'))
      // User is connected, his email address is used by default
      } else if (this.user && (!this.user.email_address || !this.regexpPattern.test(this.user.email_address.trim()))) {
        this.$awn.alert(this.$t('your_email_address_is_invalid'))
      } else if (!this.form.email || !this.regexpPattern.test(this.form.email.trim())) {
        this.$awn.alert(this.$t('recipient_email_address_is_invalid'))
      } else {
        try {
          this.form.your_email = this.form.your_email.trim()
          this.form.email = this.form.email.trim()
          const { data } = await this.form.post(process.env.API_URL + '/api/page/' + this.pages.details.id + '/share')
          console.log(this.pages.details)
          this.$awn.success(this.pages.details.name + ' - ' + this.$t(data.message))
          this.closeModal()
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
        }
      }
    },

    closeModal () {
      $('#admin-share').modal('hide')
    },

    launchModal () {
      $('#admin-share').modal('show')
    },

    destroyModal () {
      $('#admin-share').modal('dispose')
    }
  }
}
</script>
