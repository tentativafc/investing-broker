<template>
  <div class="bg-primary">
    <div class="container">
      <div class="row">
        <div class="col-12">
          <div class="d-flex align-items-center min-vh-100">
            <div class="w-100 d-block bg-white shadow-lg rounded my-5">
              <div class="row">
                <div
                  class="col-lg-5 d-none d-lg-block bg-login rounded-left"
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
                    <h1 class="h5 mb-1">Bem vindo novamente!</h1>
                    <p class="text-muted mb-4">
                      Preencha seu email e senha para acesso ao aplicativo.
                    </p>

                    <form class="user" id="app" novalidate @submit="submit">
                      <div class="form-group">
                        <input
                          type="email"
                          class="form-control"
                          :class="{ 'is-invalid': $v.email.$error }"
                          id="exampleInputEmail"
                          placeholder="Email"
                          v-model="email"
                          required
                        />
                        <div class="invalid-feedback" v-if="!$v.email.required">
                          Campo obrigatório
                        </div>
                        <div class="invalid-feedback" v-if="!$v.email.email">
                          Email inválido
                        </div>
                      </div>
                      <div class="form-group">
                        <input
                          type="password"
                          class="form-control"
                          :class="{ 'is-invalid': $v.password.$error }"
                          id="password"
                          placeholder="Senha"
                          v-model="password"
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
                      <button type="submit" class="btn btn-success btn-block">
                        Entrar
                      </button>
                    </form>

                    <div class="row mt-4">
                      <div class="col-12 text-center">
                        <p class="text-muted mb-2">
                          <a
                            href="/recover"
                            class="text-muted font-weight-medium ml-1"
                            >Esqueceu sua senha?</a
                          >
                        </p>
                        <p class="text-muted mb-0">
                          Não possui conta?
                          <a
                            href="/register"
                            class="text-muted font-weight-medium ml-1"
                            ><b>Registre-se</b></a
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
import { LOGIN } from '@/store/actions.type'
import { required, email, minLength, maxLength } from 'vuelidate/lib/validators'

export default {
  name: 'LoginComponent',
  data() {
    return {
      email: null,
      password: null
    }
  },
  validations: {
    email: {
      required,
      email
    },
    password: {
      required,
      minLength: minLength(6),
      maxLength: maxLength(20)
    }
  },
  methods: {
    submit: function(e) {
      e.preventDefault()
      this.$v.$touch()
      if (!this.$v.$invalid) {
        this.$store
          .dispatch(LOGIN, { email: this.email, password: this.password })
          .then(() => this.$router.push({ name: 'home' }))
      }
    }
  }
}
</script>
