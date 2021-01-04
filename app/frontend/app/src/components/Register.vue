<template>
  <div class="bg-primary">
    <div class="container">
      <div class="row">
        <div class="col-12">
          <div class="d-flex align-items-center min-vh-100">
            <div class="w-100 d-block bg-white shadow-lg rounded my-5">
              <div class="row">
                <div
                  class="col-lg-5 d-none d-lg-block bg-register rounded-left"
                ></div>
                <div class="col-lg-7">
                  <div class="p-5">
                    <div class="text-center">
                      <a href="index.html" class="d-block mb-5">
                        <img
                          src="assets/images/logo-dark.png"
                          alt="app-logo"
                          height="18"
                        />
                      </a>
                    </div>
                    <h1 class="h5 mb-1">Crie uma nova conta!</h1>
                    <p class="text-muted mb-4">
                      Não possui uma conta? Crie uma, leva menos de 1 minuto.
                    </p>

                    <form class="user" id="app" @submit="submit" novalidate>
                      <div class="form-group row">
                        <div class="col-sm-6 mb-3 mb-sm-0">
                          <input
                            type="text"
                            class="form-control"
                            :class="{ 'is-invalid': $v.firstname.$error }"
                            id="firstname"
                            v-model="firstname"
                            placeholder="Primeiro nome"
                            required
                          />
                          <div
                            class="invalid-feedback"
                            v-if="!$v.firstname.required"
                          >
                            Campo obrigatório
                          </div>
                        </div>
                        <div class="col-sm-6">
                          <input
                            type="text"
                            class="form-control"
                            :class="{ 'is-invalid': $v.lastname.$error }"
                            id="lastname"
                            v-model="lastname"
                            placeholder="Último nome"
                            required
                          />
                          <div
                            class="invalid-feedback"
                            v-if="!$v.lastname.required"
                          >
                            Campo obrigatório
                          </div>
                        </div>
                      </div>
                      <div class="form-group">
                        <input
                          type="email"
                          class="form-control"
                          :class="{ 'is-invalid': $v.email.$error }"
                          id="email"
                          v-model="email"
                          placeholder="Email"
                          required
                        />
                        <div class="invalid-feedback" v-if="!$v.email.required">
                          Campo obrigatório
                        </div>
                        <div class="invalid-feedback" v-if="!$v.email.email">
                          Email inválido
                        </div>
                      </div>
                      <div class="form-group row">
                        <div class="col-sm-6 mb-3 mb-sm-0">
                          <input
                            type="password"
                            class="form-control"
                            :class="{ 'is-invalid': $v.password.$error }"
                            v-model="password"
                            id="password"
                            placeholder="Senha"
                            required
                          />
                          <div
                            class="invalid-feedback"
                            v-if="!$v.password.required"
                          >
                            Campo obrigatório
                          </div>
                          <div
                            class="invalid-feedback"
                            v-if="!$v.password.minLength"
                          >
                            Campo precisa ter no mínimo
                            {{ $v.password.$params.minLength.min }} caracteres.
                          </div>
                          <div
                            class="invalid-feedback"
                            v-if="!$v.password.maxLength"
                          >
                            Campo precisa ter no máximo
                            {{ $v.password.$params.maxLength.max }} caracteres.
                          </div>
                        </div>
                        <div class="col-sm-6">
                          <input
                            type="password"
                            class="form-control"
                            :class="{ 'is-invalid': $v.repeat_password.$error }"
                            v-model="repeat_password"
                            id="repeat_password"
                            placeholder="Confirmação de senha"
                            required
                          />
                          <div
                            class="invalid-feedback"
                            v-if="!$v.repeat_password.required"
                          >
                            Campo obrigatório
                          </div>
                          <div
                            class="invalid-feedback"
                            v-if="!$v.repeat_password.minLength"
                          >
                            Campo precisa ter no mínimo
                            {{ $v.repeat_password.$params.minLength.min }}
                            caracteres.
                          </div>
                          <div
                            class="invalid-feedback"
                            v-if="!$v.repeat_password.maxLength"
                          >
                            Campo precisa ter no máximo
                            {{ $v.repeat_password.$params.maxLength.max }}
                            caracteres.
                          </div>
                          <div
                            class="invalid-feedback"
                            v-if="!$v.repeat_password.sameAs"
                          >
                            A senha e a confirmação de senha são diferentes.
                          </div>
                        </div>
                      </div>
                      <button type="submit" class="btn btn-success btn-block">
                        Criar conta
                      </button>
                    </form>

                    <div class="row mt-4">
                      <div class="col-12 text-center">
                        <p class="text-muted mb-0">
                          Já possui uma conta?
                          <a
                            href="/login"
                            class="text-muted font-weight-medium ml-1"
                            ><b>Entre</b></a
                          >
                        </p>
                      </div>
                      <!-- end col -->
                    </div>
                    <!-- end row -->
                  </div>
                  <!-- end .padding-5 -->
                </div>
                <!-- end col -->
              </div>
              <!-- end row -->
            </div>
            <!-- end .w-100 -->
          </div>
          <!-- end .d-flex -->
        </div>
        <!-- end col-->
      </div>
      <!-- end row -->
    </div>
    <!-- end container -->
  </div>
  <!-- end page -->
</template>
<script>
import {
  required,
  email,
  minLength,
  maxLength,
  sameAs
} from 'vuelidate/lib/validators'
import { REGISTER } from '@/store/actions.type'

export default {
  name: 'RegisterComponent',
  data() {
    return {
      firstname: null,
      lastname: null,
      email: null,
      password: null,
      repeat_password: null
    }
  },
  validations: {
    firstname: {
      required
    },
    lastname: {
      required
    },
    email: {
      required,
      email
    },
    password: {
      required,
      minLength: minLength(6),
      maxLength: maxLength(20)
    },
    repeat_password: {
      required,
      minLength: minLength(6),
      maxLength: maxLength(20),
      sameAsPassword: sameAs('password')
    }
  },

  methods: {
    submit: function(e) {
      e.preventDefault()

      this.$v.$touch()
      if (!this.$v.$invalid) {
        this.$store
          .dispatch(REGISTER, {
            firstname: this.firstname,
            lastname: this.lastname,
            email: this.email,
            password: this.password
          })
          .then(() => this.$router.push({ name: 'home' }))
      }
    }
  }
}
</script>
