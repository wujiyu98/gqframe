// let st = { name: "wujiy", age: 18 };

let skus = [
  { id: 1, title: "red_4G_128", image: "/sfa", price: 131, quantity: 10, group: ["red", "4G", "128"] },
  { id: 2, title: "red_4G_256", image: "/sfa", price: 131, quantity: 6, group: ["red", "4G", "256"] },
  { id: 3, title: "blue_4G_128", image: "/sfa", price: 131, quantity: 0, group: ["blue", "4G", "128"] },
  { id: 4, title: "blue_4G_256", image: "/sfa", price: 131, quantity: 6, group: ["blue", "4G", "256"] },
]
let specs = [
  { attr_name: "color", lists: ["red", "blue"] },
  { attr_name: "memory", lists: ["4G", "8G"] },
  { attr_name: "disk", lists: ["128", "256"] }
]

let a1 = ["red", "4G", "256"]
let a2 = ["red", "4G", "254"]



function getInclude(arr1, arr2) {
  let tmp = []
  tmp = arr2.filter((v) => {
    if (arr1.includes(v)) {
      return true
    }
  })
  if (tmp.length < arr2.length) {
    return false

  } else {
    return true
  }
}


function getSkuList() {
  let newSku = []
  let select_attr = ["color"]
  let select_vals = ["blue"]
  specs.forEach((spec) => {
    let newLists = []
    spec.lists.forEach((list) => {
      let disabled = true
      let cmb = []
      if (!select_attr.includes(spec.attr_name)) {
        cmb.push(...select_vals)

      }
      cmb.push(list)
      let quantity = 0
      let f_skus = skus.filter((f) => {
        if (getInclude(f.group, cmb)) {
          quantity += f.quantity
          return true
        }
      })
      if (f_skus.length > 0 && quantity > 0) {
        disabled = false
      }
      let tmp = { "attr": list, "disabled": disabled }
      newLists.push(tmp)


    })
    newSku.push({ attr_name: spec.attr_name, lists: newLists })
  })
  console.log('newSku :>> ', newSku);

}
getSkuList()





