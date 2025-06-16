{{- $global := . }}
{{- $templateID := printf "%s_%s" .Package .StructName }}
{{- if .IsAdd }}
// 
{{- range .Fields}}
    {{- if .FieldSearchType}}
{{ GenerateSearchFormItem .}}
    {{ end }}
{{ end }}


// 

{{- range .Fields}}
    {{- if .Table}}
       {{ GenerateTableColumn . }}
    {{- end }}
{{- end }}

// 
{{- range .Fields}}
   {{- if .Form}}
     {{ GenerateFormItem . }}
   {{- end }}
{{- end }}

// 

{{- range .Fields}}
              {{- if .Desc }}
    {{ GenerateDescriptionItem . }}
              {{- end }}
            {{- end }}

// 
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}

// setOptions

{{- range $index, $element := .DictTypes }}
    {{ $element }}Options.value = await getDictFunc('{{$element}}')
{{- end }}

// formData（）
{{- range .Fields}}
          {{- if .Form}}
            {{ GenerateDefaultFormValue . }}
          {{- end }}
        {{- end }}
// 

{{- range .Fields }}
        {{- if .Form }}
            {{- if eq .Require true }}
{{.FieldJson }} : [{
    required: true,
    message: '{{ .ErrorText }}',
    trigger: ['input','blur'],
},
               {{- if eq .FieldType "string" }}
{
    whitespace: true,
    message: '',
    trigger: ['input', 'blur'],
}
              {{- end }}
],
            {{- end }}
        {{- end }}
    {{- end }}



{{- if .HasDataSource }}
// 
get{{.StructName}}DataSource,

//  
const dataSource = ref({})
const getDataSourceFunc = async()=>{
  const res = await get{{.StructName}}DataSource()
  if (res.code === 0) {
    dataSource.value = res.data || []
  }
}
getDataSourceFunc()
{{- end }}

{{- else }}

{{- if not .OnlyTemplate}}
<template>
  <div>
  {{- if not .IsTree }}
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      {{- if .GvaModel }}
      <el-form-item label="" prop="createdAtRange">
      <template #label>
        <span>
          
          <el-tooltip content="（）（）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
         <el-date-picker
                  v-model="searchInfo.createdAtRange"
                  class="w-[380px]"
                  type="datetimerange"
                  range-separator=""
                  start-placeholder=""
                  end-placeholder=""
                />
       </el-form-item>
      {{ end -}}
           {{- range .Fields}} {{- if .FieldSearchType}} {{- if not .FieldSearchHide }}
            {{ GenerateSearchFormItem .}}
           {{ end }}{{ end }}{{ end }}
        <template v-if="showAllQuery">
          <!--  -->
          {{- range .Fields}}  {{- if .FieldSearchType}} {{- if .FieldSearchHide }}
          {{ GenerateSearchFormItem .}}
          {{ end }}{{ end }}{{ end }}
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"></el-button>
          <el-button icon="refresh" @click="onReset"></el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery"></el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else></el-button>
        </el-form-item>
      </el-form>
    </div>
  {{- end }}
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.add"{{ end }} type="primary" icon="plus" @click="openDialog()"></el-button>
            <el-button {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.batchDelete"{{ end }} icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete"></el-button>
            {{ if .HasExcel -}}
            <ExportTemplate {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.exportTemplate"{{ end }} template-id="{{$templateID}}" />
            <ExportExcel {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.exportExcel"{{ end }} template-id="{{$templateID}}" filterDeleted/>
            <ImportExcel {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.importExcel"{{ end }} template-id="{{$templateID}}" @on-success="getTableData" />
            {{- end }}
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="{{.PrimaryField.FieldJson}}"
        @selection-change="handleSelectionChange"
        {{- if .NeedSort}}
        @sort-change="sortChange"
        {{- end}}
        >
        <el-table-column type="selection" width="55" />
        {{ if .GvaModel }}
        <el-table-column sortable align="left" label="" prop="CreatedAt" {{- if .IsTree }} min-{{- end }}width="180">
            <template #default="scope">{{ "{{ formatDate(scope.row.CreatedAt) }}" }}</template>
        </el-table-column>
        {{ end }}
        {{- range .Fields}}
        {{- if .Table}}
            {{ GenerateTableColumn . }}
        {{- end }}
        {{- end }}
        <el-table-column align="left" label="" fixed="right" min-width="240">
            <template #default="scope">
            {{- if .IsTree }}
            <el-button {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.add"{{ end }} type="primary" link class="table-button" @click="openDialog(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon></el-button>
            {{- end }}
            <el-button {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.info"{{ end }} type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon></el-button>
            <el-button {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.edit"{{ end }} type="primary" link icon="edit" class="table-button" @click="update{{.StructName}}Func(scope.row)"></el-button>
            <el-button {{ if .IsTree }}v-if="!scope.row.children?.length"{{ end }} {{ if $global.AutoCreateBtnAuth }}v-auth="btnAuth.delete"{{ end }} type="primary" link icon="delete" @click="deleteRow(scope.row)"></el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{"{{"}}type==='create'?'':''{{"}}"}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog"> </el-button>
                  <el-button @click="closeDialog"> </el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
          {{- if .IsTree }}
            <el-form-item label=":" prop="parentID" >
                <el-tree-select
                    v-model="formData.parentID"
                    :data="[rootNode,...tableData]"
                    check-strictly
                    :render-after-expand="false"
                    :props="defaultProps"
                    clearable
                    style="width: 240px"
                    placeholder=""
                />
            </el-form-item>
          {{- end }}
        {{- range .Fields}}
          {{- if .Form}}
             {{ GenerateFormItem . }}
          {{- end }}
       {{- end }}
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="">
            <el-descriptions :column="1" border>
            {{- if .IsTree }}
            <el-descriptions-item label="">
                <el-tree-select
                  v-model="detailFrom.parentID"
                  :data="[rootNode,...tableData]"
                  check-strictly
                  disabled
                  :render-after-expand="false"
                  :props="defaultProps"
                  clearable
                  style="width: 240px"
                  placeholder=""
                />
            </el-descriptions-item>
            {{- end }}
            {{- range .Fields}}
              {{- if .Desc }}
                 {{ GenerateDescriptionItem . }}
              {{- end }}
            {{- end }}
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  {{- if .HasDataSource }}
    get{{.StructName}}DataSource,
  {{- end }}
  create{{.StructName}},
  delete{{.StructName}},
  delete{{.StructName}}ByIds,
  update{{.StructName}},
  find{{.StructName}},
  get{{.StructName}}List
} from '@/plugin/{{.Package}}/api/{{.PackageName}}'

{{- if or .HasPic .HasFile}}
import { getUrl } from '@/utils/image'
{{- end }}
{{- if .HasPic }}
// 
import SelectImage from '@/components/selectImage/selectImage.vue'
{{- end }}

{{- if .HasRichText }}
// 
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'
{{- end }}

{{- if .HasFile }}
// 
import SelectFile from '@/components/selectFile/selectFile.vue'
{{- end }}

{{- if .HasArray}}
// 
import ArrayCtrl from '@/components/arrayCtrl/arrayCtrl.vue'
{{- end }}

//  
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
{{- if .AutoCreateBtnAuth }}
// 
import { useBtnAuth } from '@/utils/btnAuth'
{{- end }}

{{if .HasExcel -}}
// 
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'
{{- end}}


defineOptions({
    name: '{{.StructName}}'
})

{{- if .AutoCreateBtnAuth }}
// 
    const btnAuth = useBtnAuth()
{{- end }}

// loading
const btnLoading = ref(false)

// /
const showAllQuery = ref(false)

// （）
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}
const formData = ref({
        {{- if .IsTree }}
            parentID:undefined,
        {{- end }}
        {{- range .Fields}}
          {{- if .Form}}
            {{ GenerateDefaultFormValue . }}
          {{- end }}
        {{- end }}
        })

{{- if .HasDataSource }}
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await get{{.StructName}}DataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()
{{- end }}



// 
const rule = reactive({
    {{- range .Fields }}
        {{- if .Form }}
            {{- if eq .Require true }}
               {{.FieldJson }} : [{
                   required: true,
                   message: '{{ .ErrorText }}',
                   trigger: ['input','blur'],
               },
               {{- if eq .FieldType "string" }}
               {
                   whitespace: true,
                   message: '',
                   trigger: ['input', 'blur'],
              }
              {{- end }}
              ],
            {{- end }}
        {{- end }}
    {{- end }}
})

const elFormRef = ref()
const elSearchFormRef = ref()

// ===========  ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

{{- if .NeedSort}}
// 
const sortChange = ({ prop, order }) => {
  const sortMap = {
    CreatedAt:"CreatedAt",
    ID:"ID",
    {{- range .Fields}}
     {{- if .Table}}
      {{- if and .Sort}}
        {{- if not (eq .ColumnName "")}}
            {{.FieldJson}}: '{{.ColumnName}}',
        {{- end}}
      {{- end}}
     {{- end}}
    {{- end}}
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}
{{- end}}

{{- if not .IsTree }}
// 
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    {{- range .Fields}}{{- if eq .FieldType "bool" }}
    if (searchInfo.value.{{.FieldJson}} === ""){
        searchInfo.value.{{.FieldJson}}=null
    }{{ end }}{{ end }}
    getTableData()
  })
}

// 
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 
const getTableData = async() => {
  const table = await get{{.StructName}}List({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
{{- else }}
// 
const defaultProps = {
  children: "children",
  label: "{{ .TreeJson }}",
  value: "{{ .PrimaryField.FieldJson }}"
}

const rootNode = {
  {{ .PrimaryField.FieldJson }}: 0,
  {{ .TreeJson }}: '',
  children: []
}

// 
const getTableData = async() => {
  const table = await get{{.StructName}}List()
  if (table.code === 0) {
    tableData.value = table.data || []
  }
}
{{- end }}

getTableData()

// ==============  ===============

//   
const setOptions = async () =>{
{{- range $index, $element := .DictTypes }}
    {{ $element }}Options.value = await getDictFunc('{{$element}}')
{{- end }}
}

//   
setOptions()


// 
const multipleSelection = ref([])
// 
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 
const deleteRow = (row) => {
    ElMessageBox.confirm('?', '', {
        confirmButtonText: '',
        cancelButtonText: '',
        type: 'warning'
    }).then(() => {
            delete{{.StructName}}Func(row)
        })
    }

// 
const onDelete = async() => {
  ElMessageBox.confirm('?', '', {
    confirmButtonText: '',
    cancelButtonText: '',
    type: 'warning'
  }).then(async() => {
      const {{.PrimaryField.FieldJson}}s = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: ''
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          {{.PrimaryField.FieldJson}}s.push(item.{{.PrimaryField.FieldJson}})
        })
      const res = await delete{{.StructName}}ByIds({ {{.PrimaryField.FieldJson}}s })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: ''
        })
        if (tableData.value.length === {{.PrimaryField.FieldJson}}s.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// （）
const type = ref('')

// 
const update{{.StructName}}Func = async(row) => {
    const res = await find{{.StructName}}({ {{.PrimaryField.FieldJson}}: row.{{.PrimaryField.FieldJson}} })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 
const delete{{.StructName}}Func = async (row) => {
    const res = await delete{{.StructName}}({ {{.PrimaryField.FieldJson}}: row.{{.PrimaryField.FieldJson}} })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: ''
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 
const dialogFormVisible = ref(false)

// 
const openDialog = ({{- if .IsTree -}}row{{- end -}}) => {
    type.value = 'create'
    {{- if .IsTree }}
    formData.value.parentID = row ? row.{{.PrimaryField.FieldJson}} : undefined
    {{- end }}
    dialogFormVisible.value = true
}

// 
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
    {{- range .Fields}}
      {{- if .Form}}
        {{ GenerateDefaultFormValue . }}
      {{- end }}
    {{- end }}
        }
}
// 
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await create{{.StructName}}(formData.value)
                  break
                case 'update':
                  res = await update{{.StructName}}(formData.value)
                  break
                default:
                  res = await create{{.StructName}}(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '/'
                })
                closeDialog()
                getTableData()
              }
      })
}

const detailFrom = ref({})

// 
const detailShow = ref(false)


// 
const openDetailShow = () => {
  detailShow.value = true
}


// 
const getDetails = async (row) => {
  // 
  const res = await find{{.StructName}}({ {{.PrimaryField.FieldJson}}: row.{{.PrimaryField.FieldJson}} })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>
{{if .HasFile }}
.file-list{
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.fileBtn{
  margin-bottom: 10px;
}

.fileBtn:last-child{
  margin-bottom: 0;
}
{{end}}
</style>
{{- else}}
<template>
<div>form</div>
</template>
<script setup>
defineOptions({
  name: '{{.StructName}}'
})
</script>
<style>
</style>
{{- end }}

{{- end }}
