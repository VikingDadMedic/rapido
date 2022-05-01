<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>, Félix P <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<!-- MODAL TO CREATE NEW CONTACT FORM -->
<template>
  <Modal id="admin-newcontactform" :title="$t('settings_contents__new_contact_form')" customClass="modal fade in">
    <form @submit.prevent="saveContactForm" @keydown="form.onKeydown($event)">

      <!-- Contact form name -->
      <div class="form-group row">
        <label class="col-md-4 col-form-label text-md-right">
          {{ $t('settings_contents__contact_form_name') }}
        </label>
        <span class="col-md-8">
          <input v-model="form.name" type="text" name="name" class="form-control">
        </span>
      </div>

      <!-- Recipient email address -->
      <div class="form-group row">
        <label class="col-md-4 col-form-label text-md-right">
          {{ $t('recipient_email_address') }}
        </label>
        <span class="col-md-8">
          <input v-model="form.json_settings.to.value" type="text" name="to" class="form-control">
        </span>
      </div>

    </form>

      <!-- Create button -->
      <div class="form-group row">
        <span class="col-md-8 offset-md-4 d-flex justify-content-start">
          <v-button :loading="form.busy" slot="footer" @click="saveContactForm(form)" class="btn btn-success">{{ $t('button__create') }}</v-button>
        </span>
      </div>

  </Modal>
</template>

<script>
import { mapGetters } from 'vuex'
import Form from 'vform'
import axios from 'axios'
import $ from 'jquery'
export default {
  name: 'admin-newcontactform',

  mounted () {
    this.setForm()
    this.$nextTick(() => {
      if (this.$route.query.new) {
        this.launchModal()
      }
    })
  },

  data () {
    return {
      form: new Form({
        id: 2,
        access_level: 0,
        name: '',
        // created_by: null,
        extension: '',
        // Form json_settings obj struct
        json_settings: {
          to: {
            name: 'to',
            placeholder: 'Email address to send messages to',
            type: 'text',
            value: '',
            regexpPattern: /^.+@.+$/
          },

          layout: {
            name: 'layout',
            placeholder: 'Email template',
            type: 'textarea',
            value:
              '\u003ch1\u003eSomeone with email {{ .email }}\u003c/h1\u003e\u003cp\u003eThe message is\u003c/p\u003e\u003cp\u003e{{ .message }}.\u003c/p\u003e'
          }
        }
      }),
      // Json_settings array struct
      json_settings: [
        {
          name: 'to',
          placeholder: 'Email address to send messages to',
          type: 'text',
          value: ''
        },
        {
          name: 'layout',
          placeholder: 'Email template',
          type: 'textarea',
          value:
            '\u003ch1\u003eSomeone with email {{ .email }}\u003c/h1\u003e\u003cp\u003eThe message is\u003c/p\u003e\u003cp\u003e{{ .message }}.\u003c/p\u003e'
        }
      ]
    }
  },

  computed: mapGetters({
    access_level: 'page/access_level',
    user: 'auth/user'
  }),

  beforeDestroy () {
    // destroy modal on component destroy
    this.destroyModal()
  },

  methods: {
    // Init new form with default values
    initForm () {
      this.form = new Form({
        access_level: 0,
        name: '',
        json_settings: {
          to: {
            name: 'to',
            placeholder: 'Email address to send messages to',
            type: 'text',
            value: this.user.email_address,
            regexpPattern: /^.+@.+$/
          },

          layout: {
            name: 'layout',
            placeholder: 'Email template',
            type: 'textarea',
            value:
              '\u003ch1\u003eSomeone with email {{ .email }}\u003c/h1\u003e\u003cp\u003eThe message is\u003c/p\u003e\u003cp\u003e{{ .message }}.\u003c/p\u003e'
          }
        } // ,
        // created_by: this.user.id
      })
    },

    // Save new contactForm to database
    async saveContactForm (v) {
      if (!v.name || v.name.trim() === '') {
        this.$awn.alert(this.$t('settings_contents__empty_form_name'))
      } else if (!v.json_settings.to.value || !v.json_settings.to.regexpPattern.test(v.json_settings.to.value.trim())) {
        this.$awn.alert(this.$t('recipient_email_address_is_invalid'))
      } else {
        v.name = v.name.trim()
        v.json_settings.to.value = v.json_settings.to.value.trim()
        let obj = JSON.parse(JSON.stringify(v))
        // use json_settings array
        this.json_settings[0].value = v.json_settings.to.value
        obj.json_settings = this.json_settings
        try {
          await axios.post(process.env.API_URL + '/api/contactform', obj)
          this.$awn.success(this.$t('settings_contents__contact_form_created'))
          await this.$store.dispatch('page/fetchAllContactForm')
          this.closeModal()
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message))
        }
      }
    },

    // Set certain form values
    setForm () {
      console.log(this.user.email_address)
      // this.form['created_by'] = this.user.id
      this.form.json_settings.to.value = this.user.email_address
    },
    // Fetch all contact form to refresh displayed data if data changed
    async fetchAllContactForm () {
      await this.$store.dispatch('page/fetchAllContactForm')
    },

    // Launch modal
    launchModal () {
      this.initForm()
      $('#admin-newcontactform').modal('show')
    },
    // Close modal
    closeModal () {
      $('#admin-newcontactform').modal('hide')
    },
    // Destroy modal
    destroyModal () {
      $('#admin-newcontactform').modal('dispose')
    }
  }
}
</script>
