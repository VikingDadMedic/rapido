// Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
// Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
// Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

// This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.


/* eslint-disable */
/**
 * Custom module for quilljs to allow user to drag images from their file system into the editor
 * and paste images from clipboard (Works on Chrome, Firefox, Edge, not on Safari)
 * @see https://quilljs.com/blog/building-a-custom-module/
 */
import axios from 'axios'
import store from '@/store'
import $ from 'jquery'
import { EventBus } from './event-bus'
export class ImageUpload {

	/**
	 * Instantiate the module given a quill instance and any options
	 * @param {Quill} quill
	 * @param {Object} options
	 */
	constructor(quill, options = {}) {
		// save the quill reference
		this.quill = quill;
		// save options
    this.options = options;
		// listen for drop and paste events
		this.quill.getModule('toolbar').addHandler('customimage', this.selectLocalImage.bind(this));
    // this.quill.getModule('toolbar').addHandler('customimage',
    //   (value) => {
    //     if (value === 1) {
    //       vm.selectLocalImage.bind(vm)
    //     }
    //   }
    // );
	}


	/**
	 * Select local image
	 */
	selectLocalImage(v) {
    if (v === '0') return
    if (v === '1') {
      const input = document.createElement('input');
      input.setAttribute('type', 'file');
      input.click();

      // Listen upload local image and save to server
      input.onchange = () => {
        const file = input.files[0];

        // file type is only image.
        if (/^image\//.test(file.type)) {
          const checkBeforeSend = this.options.checkBeforeSend || this.checkBeforeSend.bind(this);
          checkBeforeSend(file, this.sendToServer.bind(this));
        } else {
          console.warn('You could only upload images.');
        }
      };
    } else {
      const selectFromGallery = this.options.selectFromGallery || this.selectFromGallery.bind(this);
      selectFromGallery()
    }
  }


  /**
   * Select from gallery
   */
  selectFromGallery() {
    $('#gallery-modal').modal('show')
    const callbackOK = this.options.callbackOK || this.uploadImageCallbackOK.bind(this)
    let vm = this
    EventBus.$on('my-event', (data) => {
      callbackOK(data.url, vm.insertURL.bind(vm));
      // console.log('Inside `my-event`. It\'ll be executed only one time!', data);
    });
  }


	/**
	 * Check file before sending to the server
	 * @param {File} file
	 * @param {Function} next
	 */
	checkBeforeSend(file, next) {
		next(file);
	}

	/**
	 * Send to server
	 * @param {File} file
	 */
	async sendToServer(file) {
		const url = this.options.url || 'your-url.com',
			method = this.options.method || 'POST',
			headers = this.options.headers || {},
			callbackOK = this.options.callbackOK || this.uploadImageCallbackOK.bind(this),
			callbackKO = this.options.callbackKO || this.uploadImageCallbackKO.bind(this);

		// const xhr = new XMLHttpRequest();
		// // init http query
		// xhr.open(method, url, true);
		// // add custom headers
		// for (var index in headers) {
		// 	xhr.setRequestHeader(index, headers[index]);
		// }

		// // listen callback
		// xhr.onload = () => {
		// 	if (xhr.status === 200) {
		// 		callbackOK(JSON.parse(xhr.responseText), this.insert.bind(this));
		// 	} else {
		// 		callbackKO({
		// 			code: xhr.status,
		// 			type: xhr.statusText,
		// 			body: xhr.responseText
		// 		});
		// 	}
		// };

		// xhr.send(fd);
    const token = store.getters['auth/token']
    headers.Authorization = `Bearer ${token}`
    try {
      let files = new FormData()
      files.append('files', file)
      let { data } = await axios.post(url, files, { headers: headers })
      callbackOK(data, this.insert.bind(this));
    } catch (e) {
  		callbackKO({
  			code: 400,
  			type: 'error',
  			body: 'Failed to upload file'
  		});
    }
	}

	/**
	 * Insert the image into the document at the current cursor position
	 * @param {String} dataUrl  The base64-encoded image URI
	 */
	insert(dataUrl) {
    let url = window.location.href
    let arr = url.split("/")
    let prot = "//"
		const index = (this.quill.getSelection() || {}).index || this.quill.getLength();
		this.quill.insertEmbed(index, 'image', dataUrl.success[0].url.includes('://') ? dataUrl.success[0].url : prot + dataUrl.success[0].url, 'user');
	}

	/**
	 * Insert the image into the document at the current cursor position
	 * @param {String} dataUrl  The base64-encoded image URI
	 */
	insertURL(dataUrl) {
    let prot = "//"
		const index = (this.quill.getSelection() || {}).index || this.quill.getLength();
		this.quill.insertEmbed(index, 'image', dataUrl.includes('://') ? dataUrl : prot + dataUrl, 'user');
	}

	/**
	 * callback on image upload succesfull
	 * @param {Any} response http response
	 */
	uploadImageCallbackOK(response, next) {
		next(response);
	}

	/**
	 * callback on image upload failed
	 * @param {Any} error http error
	 */
	uploadImageCallbackKO(error) {
		alert(JSON.stringify(error));
	}


}
