$(function() {
    let tabBtn1 = $('#tab-btn-stable');
    let dropdown1 = $('#dropdown-menu-model');
    let submit = $('#submit')
    let textInput = $('#text-input')
    let imageDisplay = $('#image-display')
    let loading = $('#loading')
    let providerInput = $('#provider-input')

    /* Load Model List */
    $.get('/get/model-list', function(response) {
        $.each(response, function(_, model) {
            let row = "<option value='" + model + "'>" + model + "</option>";

            dropdown1.append(row);
        });
    });

    /* 상단 Tab Button -> Stable Diffusion으로 이동 */
    $(document).ready(function() {
        tabBtn1.click(function() {
            $.get('/get/url/stable_diffusion', function(response) {
                url = response.url;

                window.location.href = url;
            });
        });
    });

    /* Image Inference 요청 */
    $(document).ready(function() {
        submit.click(function() {
            //inference();
            //inferenceTest();
            inferenceTest1();
        });
    });

    /* ========== default ========== */
    function inference() {
        let prompt = textInput.val();
        let model = dropdown1.val();
        let provider = providerInput.val()
        let data = {
            prompt: prompt,
            provider: provider
        };

        if (model == "Select a model...") {
            alert("Please select a model");

            return;
        } else if (!prompt.trim()) {
            alert("Please enter a prompt.");

            return;
        }

        imageDisplay.empty();
        loading.text("Loading...")

        fetch("/model/" + model + "/infer", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(response => response.json())
        .then(data => {
            let base64Image = data.image;
            let imgTag = $("<img>", { src: "data:image/png;base64," + base64Image, alt: "Generated Image" });
            
            imageDisplay.empty().append(imgTag);

           loading.text("Finish");
       })
        .catch(error => {
            console.error("ERROR" + error);

            loading.text("Error")
        });
    }

    /* ========== Test ========== */
    function inferenceTest() {
        let numberOfImages = parseInt(prompt("How many images do you want to generate?"));
            if (!isNaN(numberOfImages) && numberOfImages > 0) {
                generateImages(numberOfImages);
            } else {
                alert("Please enter a valid number.");
            }
    }

    function generateImages(count) {
        if (count <= 0) {
            loading.text("Finished generating images.");
            return;
        }
    
        let promptText = textInput.val();
        let model = dropdown1.val();
        let provider = providerInput.val()
        let data = {
            prompt: promptText,
            provider: provider
        };
    
        if (model == "Select a model...") {
            alert("Please select a model");
            return;
        } else if (!promptText.trim()) {
            alert("Please enter a prompt.");
            return;
        }
    
        fetchImage(model, data, function() {
            generateImages(count - 1);
        });
    }
    
    function fetchImage(model, data, callback) {
        loading.text("Loading...");
    
        fetch("/model/" + model + "/infer", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(response => response.json())
        .then(data => {
            let base64Image = data.image;
            let imgTag = $("<img>", { src: "data:image/png;base64," + base64Image, alt: "Generated Image" });
            
            imageDisplay.empty().append(imgTag);


            loading.text("");
            if (callback) {
                callback();
            }
        })
        .catch(error => {
            console.error("ERROR" + error);
            loading.text("Error");
            if (callback) {
                callback();
            }
        });
    }

    /* ========== Test 1 ========== */
    function inferenceTest1() {
        let numberOfImages = parseInt(prompt("How many images do you want to generate?"));
        if (!isNaN(numberOfImages) && numberOfImages > 0) {
            let startTime = performance.now();
            loading.empty();
            imageDisplay.empty();
            
            generateImages1(numberOfImages, startTime);
        } else {
            alert("Please enter a valid number.");
        }
    }
    
    function generateImages1(count, startTime) {
        if (count <= 0) {
            let endTime = performance.now();
            let totalTime = endTime - startTime;
            loading.text("Finished generating images. Total time: " + totalTime.toFixed(2) + " milliseconds");
            return;
        }
    
        let promptText = textInput.val();
        let model = dropdown1.val();
        let provider = providerInput.val()
        let data = {
            prompt: promptText,
            provider: provider
        };
    
        if (model == "Select a model...") {
            alert("Please select a model");
            return;
        } else if (!promptText.trim()) {
            alert("Please enter a prompt.");
            return;
        }

        let promises = [];
        for (let i = 0; i < count; i++) {
            promises.push(fetchImage1(model, data));
        }
    
        Promise.all(promises)
            .then(() => {
                let endTime = performance.now();
                let totalTime = endTime - startTime;
                loading.text("Finished generating images. Total time: " + (totalTime/1000).toFixed(2) + " seconds");
            })
            .catch(error => {
                console.error("ERROR" + error);
                loading.text("Error");
            });
    }
    
    function fetchImage1(model, data) {
        return new Promise((resolve, reject) => {
            fetch("/model/" + model + "/infer", {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(data)
                })
                .then(response => response.json())
                .then(data => {
                    let base64Image = data.image;
                    let imgTag = $("<img>", {
                        src: "data:image/png;base64," + base64Image,
                        alt: "Generated Image"
                    });
    
                    imageDisplay.append(imgTag);
                    resolve();
                })
                .catch(error => {
                    console.error("ERROR" + error);
                    reject(error);
                });
        });
    }
});